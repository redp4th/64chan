package middleware

import (
	"net/http"
	"nmb/config"
	"nmb/pkg/service"
	"nmb/pkg/util"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data interface{}
		code := 0

		token := c.Request.Header.Get("Authorization")
		if token == "" {
			code = service.Unauthorized
		} else {
			claims, err := util.ParseToken(token, config.JWTSecret)
			if err != nil {
				code = service.AuthenticationFailed
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = service.TokenExpired
			}
		}

		if code != 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": service.ErrorMessage(code),
				"data":    data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
