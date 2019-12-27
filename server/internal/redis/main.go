package redis

import (
	"github.com/fedoratipper/bitkiosk/server/pkg/utils"
	"github.com/go-redis/redis/v7"
	"strconv"
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
		Password: "",
		DB: redisDB,
	})

	_, err := client.Ping().Result()

	return client,	err
}
