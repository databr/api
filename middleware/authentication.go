package middleware

import (
	"github.com/databr/api/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/status/pingdom" {
			c.Next()
			return
		}

		token, err := jwt.ParseFromRequest(c.Request, func(token *jwt.Token) (interface{}, error) {
			return config.PrivateKey, nil
		})

		if err == nil && token.Valid {

			c.Set("app_id", token.Claims["app_id"])
			c.Next()
		} else {
			c.JSON(401, map[string]interface{}{
				"error":        "token invalid",
				"message":      "Token is invalid, check your request and try again",
				"server_error": err.Error(),
			})
			c.Abort(401)
		}
	}
}
