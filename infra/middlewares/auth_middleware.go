package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/pepedoni/go-clean-arch-org-user-api/utils/errors/rest_errors"
)

const (
	SECRET_KEY_JWT = "SECRET_KEY_JWT"
)

var (
	secretKey = []byte(os.Getenv(SECRET_KEY_JWT))
)

func OAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestAccessToken string
		if len(c.Request.Header["Authorization"]) > 0 {
			requestAccessToken = c.Request.Header["Authorization"][0]
		}
		if requestAccessToken == "" {
			c.JSON(http.StatusUnauthorized, rest_errors.NewUnauthorizedError("invalid token"))
			c.Abort()
			return
		}

		token, err := jwt.Parse(requestAccessToken, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, rest_errors.NewUnauthorizedError("invalid token"))
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, rest_errors.NewUnauthorizedError("invalid token"))
			c.Abort()
			return
		}

		c.Next()
	}
}
