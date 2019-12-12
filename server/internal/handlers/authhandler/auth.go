package authhandler

import (
	"context"
	"fmt"
	"github.com/fedoratipper/bitkiosk/server/internal/logger"
	"github.com/gin-gonic/gin"
)

const (
	NoAuth        = 1 << iota
	UserAuth      = 2 << iota
	ModeratorAuth = 3 << iota
	AdminAuth     = 4 << iota
	TokenKey = "token"
);


func AuthenticationHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)

		jwt, _ := c.Cookie("REVEL_FLASH")
		c.Set(TokenKey, jwt)

		c.Next()
	}
}

func GetAuthLevelFromContext(ctx context.Context) (uint, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		logger.Warn(err)
		return NoAuth, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		logger.Warn(err)
		return NoAuth, err
	}

	token, _ :=  gc.Get(TokenKey)

	if token == "123" {
		return UserAuth, nil
	}

	return NoAuth, nil
}