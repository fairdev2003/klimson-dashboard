package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zgierz/klimson/backend/logger"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")

		if err != nil {
			logger.ErrorLog("Brak ciasteczka 'token'!")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Brak autoryzacji"})
			return
		}
		logger.GreenServerLog(tokenString)

		secret := os.Getenv("TOKEN")
		if secret == "" {
			logger.ErrorLog("Brak TOKEN w env")
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Konfiguracja serwera błąd"})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("nieoczekiwana metoda podpisu")
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			logger.ErrorLog(fmt.Sprintf("Błąd JWT: %v", err))
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Niepoprawny token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Błędne claims"})
			return
		}

		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) < time.Now().Unix() {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token wygasł"})
				return
			}
		}

		contributor, _ := claims["contributor"].(bool)
		if !contributor {
			c.Set("isRoot", true)
		} else {
			permissions, _ := claims["permissions"].(string)
			c.Set("permissions", permissions)
		}

		c.Next()
	}
}
