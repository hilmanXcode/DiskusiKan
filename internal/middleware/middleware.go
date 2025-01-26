package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hilmanXcode/DiskusiKan/internal/configs"
	"github.com/hilmanXcode/DiskusiKan/pkg/jwt"
)

func AuthMiddleWare() gin.HandlerFunc {
	secretKey := configs.Get().Service.SecretJwt

	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")

		header = strings.TrimSpace(header)

		if header == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("Invalid token"))
			return
		}

		userID, username, err := jwt.ValidateToken(header, secretKey)

		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		c.Set("userID", userID)
		c.Set("username", username)
		c.Next()
	}

}
