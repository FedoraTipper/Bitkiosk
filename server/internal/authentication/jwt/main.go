package jwt

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fedoratipper/bitkiosk/server/internal/logger"
	"github.com/fedoratipper/bitkiosk/server/pkg/utils"
	"time"
)

var signingKey []byte

func init() {
	signingKey = []byte(utils.MustGet("JWT_SIGNING_KEY"))
}

func GenerateJWT(ttl time.Duration, authLevel int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["auth_level"] = authLevel
	claims["exp"] = time.Now().Add(time.Minute * ttl).Unix()

	tokenString, err := token.SignedString(signingKey)

	if err != nil {
		logger.Error("Unable to sign token \n" + err.Error())
		return "", err
	}

	return tokenString, nil
}


func GetToken(tokenString string) *jwt.Token{
	token, err := jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Error")
		}
		return signingKey, nil
	})

	if err != nil {
		fmt.Errorf(err.Error())
		return nil
	}

	if token.Valid {
		return token
	}

	return nil
}