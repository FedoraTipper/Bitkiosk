package session

import (
	"encoding/json"
	"github.com/fedoratipper/bitkiosk/server/internal/redis"
	"github.com/fedoratipper/bitkiosk/server/pkg/utils"
	"github.com/gofrs/uuid"
	"time"
)

var signingKey []byte

type AuthLevel struct {
	AuthLevel int
	UID       int
}

const (
	NoAuth = iota
	UserAuth
	ModeratorAuth
	AdminAuth
)

func init() {
	signingKey = []byte(utils.MustGet("JWT_SIGNING_KEY"))
}

func GenerateSession (ttl time.Duration, authLevel AuthLevel) (string, error) {
	var generatedKey = ""

	redisClient, err := redis.Factory()

	if err == nil {
		generatedUUID, err := uuid.NewV4()

		if err == nil {
			generatedKey = generatedUUID.String()
			jsonPayload, err := json.Marshal(authLevel)

			if err == nil {
				err = redisClient.Set(generatedKey, jsonPayload, ttl).Err()
			}
		}
	}

	_ = redis.CloseRedisSession(redisClient)

	return generatedKey, err
}

func DestroySession (sessionId string) error {

	redisClient, err := redis.Factory()

	if err == nil {
		err = redisClient.Del(sessionId).Err()
	}

	_ = redis.CloseRedisSession(redisClient)

	return err
}


func GetSessionAuthLevel (sessionKey string) (AuthLevel, error) {
	var sessionAuthLevel AuthLevel

	redisClient, err := redis.Factory()

	if err == nil {
		keyValue, err := redisClient.Get(sessionKey).Result()

		if err == nil {
			err = json.Unmarshal([]byte(keyValue), &sessionAuthLevel)
		}
	}

	_ = redis.CloseRedisSession(redisClient)

	return sessionAuthLevel, err
}

