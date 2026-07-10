package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/zgierz/klimson/backend/logger"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			logger.ErrorLog("Nie ma nagłówka Bearer!")
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Nie ma nagłówka Bearer!"})
			return
		}

		err := godotenv.Load("../.env")
		if err != nil {
			logger.ErrorLog("Nie udalo sie zaladowac pliku .env")
			return
		}

		secret := os.Getenv("TOKEN")

		if secret == "" {
			logger.ErrorLog("Brak TOKEN w env")
			c.AbortWithStatus(http.StatusInternalServerError)
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unauthorized")
			}
			return []byte(secret), nil
		})
		if err != nil || token == nil {
			logger.ErrorLog(fmt.Sprintf("Błąd JWT: %v", err))
			c.AbortWithStatus(http.StatusUnauthorized)
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		expRaw, exists := claims["exp"]
		if !exists {
			c.Abort()
			return
		}

		expFloat, ok := expRaw.(float64)
		if !ok {
			c.Abort()
			return
		}

		exp := int64(expFloat)
		if exp < time.Now().Unix() {
			c.Abort()
			return
		}
		contributor, _ := claims["contributor"].(bool)
		if contributor != true {
			c.Set("isRoot", true)
			c.Next()
			return
		} else {
			permissions, _ := claims["permissions"].(string)
			c.Set("permissions", permissions)
			c.Next()
		}

	}
}
