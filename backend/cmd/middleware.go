package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zgierz/klimson/backend/khttp"
	"github.com/zgierz/klimson/backend/logger"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := ctx.Cookie("k-token")

		if err != nil {
			logger.ErrorLog("No cooklie 'k-token'!")
			khttp.UnauthorizedResponse(ctx, nil, "Session is unauthorized!")
			return
		}

		secret := os.Getenv("TOKEN")
		if secret == "" {
			logger.ErrorLog("Brak TOKEN w env")
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Konfiguracja serwera błąd"})
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
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Niepoprawny token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Błędne claims"})
			return
		}

		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) < time.Now().Unix() {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token wygasł"})
				return
			}
		}

		contributor, _ := claims["contributor"].(bool)
		if !contributor {
			ctx.Set("isRoot", true)
			ctx.Set("claims", claims)
		} else {
			permissions, _ := claims["permissions"].(string)
			ctx.Set("permissions", permissions)
		}

		ctx.Next()
	}
}
