package authentication

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	jwtHandler "github.com/fedoratipper/bitkiosk/server/internal/authentication/jwt"
	"github.com/fedoratipper/bitkiosk/server/internal/logger"
	"github.com/gin-gonic/gin"
)

type AuthLevel struct {
	AuthLevel uint
	UID       uint
}

const (
	NoAuth = iota
	UserAuth
	ModeratorAuth
	AdminAuth
)

const (
	tokenKey = "Token"
	authKey  = "Authorization"
)

func GetAuthLevelFromJWT(ctx context.Context) (*AuthLevel, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		logger.Warn(err)
		return &AuthLevel{NoAuth, 0}, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		logger.Warn(err)
		return &AuthLevel{NoAuth, 0}, err
	}

	token := gc.Request.Header.Get(tokenKey)

	if token != "" {
		jwtToken := jwtHandler.GetToken(fmt.Sprint(token))

		if jwtToken != nil {
			if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok {
				if claims["authlevel"] == true {
					return &AuthLevel{UserAuth, 1}, nil
				}
			}
		}
	}

	return &AuthLevel{NoAuth, 0}, nil
}

