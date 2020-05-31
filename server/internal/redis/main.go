package redis

import (
	"encoding/json"
	"github.com/fedoratipper/bitkiosk/server/internal/logger"
	"github.com/fedoratipper/bitkiosk/server/pkg/utils"
	"github.com/go-redis/redis/v7"
	"reflect"
	"strconv"
	"time"
)


var redisHost, redisPassword string
var redisPort, redisDB int

func init() {
	redisHost = utils.MustGet("REDIS_HOST")
	redisPort = utils.MustGetInt("REDIS_PORT")
	redisDB = utils.MustGetInt("REDIS_DB")
	redisPassword = utils.MustGet("REDIS_PASSWORD")
}

func Factory() (*redis.Client, error) {

	client := redis.NewClient(&redis.Options{
		Addr: redisHost + ":" + strconv.Itoa(redisPort),
		Password: redisPassword,
		DB: redisDB,
	})

	_, err := client.Ping().Result()

	if err != nil {
		logger.Errorfn("Factory", err)
	}

	return client,	err
}

func LoadObjectFromCache(objectStruct interface{}, id string) interface{}{
	redisClient, err := Factory()
	reflectedObject := reflect.TypeOf(objectStruct)

	if err == nil && redisClient != nil {
		keyValue, err := redisClient.Get(reflectedObject.Name() + ":" + id).Result()
		if err == nil && keyValue != "" {
			err = json.Unmarshal([]byte(keyValue), &objectStruct)
		}

		_ = CloseRedisSession(redisClient)
	}

	return objectStruct
}

func PutObjectInCache(object interface{}, id string) error {
	redisClient, err := Factory()

	if err == nil && redisClient != nil {
		valueToPut, err := json.Marshal(object)
		ttl, err := time.ParseDuration("50m")

		if err == nil {
			err = redisClient.Set(reflect.TypeOf(object).Name() + ":" + id, valueToPut, ttl).Err()
		}

		if err != nil {
			logger.Errorfn("PutObjectInCache", err)
		}

		_ = CloseRedisSession(redisClient)
	}

	return err
}

func CloseRedisSession(client *redis.Client) (closedCorrectly bool) {
	err := client.Close()

	if err != nil {
		logger.Errorfn("closeSession", err)
		return false
	}

	return true
}
