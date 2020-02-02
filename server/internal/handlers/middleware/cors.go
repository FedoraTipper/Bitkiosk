package middleware

import (
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "time"
)

func AddCorsToGinRequest() gin.HandlerFunc {
    config := cors.Config{
       AllowAllOrigins:  false,
       AllowOrigins:     []string{"http://localhost:8080", "http://localhost:7777"},
       AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
       AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token", "set-cookie", "Authorization"},
       AllowCredentials: true,
       MaxAge:           12 * time.Hour,
       ExposeHeaders:    []string{"Access-Token", "Uid", "Authorization"},
    }
    return cors.New(config)
}

//AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Access-Token", "Set-Cookie", "Authorization"},
