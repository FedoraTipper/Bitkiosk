package authentication

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	jwtHandler "github.com/fedoratipper/bitkiosk/server/internal/authentication/jwt"
	"github.com/fedoratipper/bitkiosk/server/internal/authentication/session"
	"github.com/fedoratipper/bitkiosk/server/internal/logger"
	"github.com/gin-gonic/gin"
)


const (
	tokenKey = "Token"
	authKey  = "Authorization"
)

// Not used - Moved to session management
func GetAuthLevelFromJWT(ctx context.Context) (*session.AuthLevel, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		logger.Warn(err)
		return &session.AuthLevel{AuthLevel: session.NoAuth}, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		logger.Warn(err)
		return &session.AuthLevel{AuthLevel: session.NoAuth}, err
	}

	token := gc.Request.Header.Get(tokenKey)

	if token != "" {
		jwtToken := jwtHandler.GetToken(fmt.Sprint(token))

		if jwtToken != nil {
			if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok {
				if claims["authlevel"] == true {
					return &session.AuthLevel{AuthLevel: session.UserAuth, UID: 1}, nil
				}
			}
		}
	}

	return &session.AuthLevel{AuthLevel: session.NoAuth}, nil
}

func GetAuthLevelFromSession(ctx context.Context) (*session.AuthLevel, error) {
	ginContext := ctx.Value("GinContextKey")

	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		logger.Warn(err)
		return &session.AuthLevel{AuthLevel: session.NoAuth}, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		logger.Warn(err)
		return &session.AuthLevel{AuthLevel: session.NoAuth}, err
	}

	authCookie, err := gc.Request.Cookie(authKey)

	if err == nil && authCookie != nil && authCookie.Value != "" {
		authLevel, _ := session.GetSessionAuthLevel(authCookie.Value)
		return &session.AuthLevel{AuthLevel: authLevel.AuthLevel, UID: authLevel.UID}, nil
	}

	return &session.AuthLevel{AuthLevel: session.NoAuth}, nil
}