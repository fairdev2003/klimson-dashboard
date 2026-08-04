package handlers

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/zgierz/klimson/backend/api"
	"github.com/zgierz/klimson/backend/logger"
	"github.com/zgierz/klimson/backend/models"
	"golang.org/x/crypto/bcrypt"
)

type RootJwtClaims struct {
	Exp   time.Time
	Name  string
	Login string
}

func GenerateRootToken(isContributor bool) (string, error) {

	claims := jwt.MapClaims{
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"name":  "Jakub Klimkiewicz",
		"login": "root",
	}

	err := godotenv.Load("../.env")
	if err != nil {
		logger.ErrorLog("Nie udalo sie zaladowac pliku .env")
	}

	secret := os.Getenv("TOKEN")
	logger.ServerLog(secret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func GenerateToken(contributor models.Contributor) (string, error) {

	claims := jwt.MapClaims{
		"exp":         time.Now().Add(time.Hour * 2).Unix(),
		"id":          contributor.ID,
		"permissions": contributor.Permissions,
		"contributor": true,
		"name":        contributor.Name,
		"login":       contributor.Login,
	}

	err := godotenv.Load("../.env")
	if err != nil {
		logger.ErrorLog("Erorr during .env preloading")
		return "", err
	}

	secret := os.Getenv("TOKEN")
	logger.ServerLog(secret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func PasswordMiddleware(hashedPassword string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pass := ctx.Query("password")

		if pass == "" {
			currentPath := ctx.Request.URL.Path
			exampleURL := fmt.Sprintf("http://%s%s?password=PASSWORD", ctx.Request.Host, currentPath)

			api.UnauthorizedResponse(ctx, map[string]interface{}{
				"error":   true,
				"message": "Password is required.",
				"hint":    "Please provide the password using the 'password' query parameter. or header 'X-Password'.",
				"example": exampleURL,
			}, "Password is required.")
			ctx.Abort()
			return
		}

		err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(pass))
		if err != nil {
			api.UnauthorizedResponse(ctx, nil, fmt.Sprintf("Password is incorrect. %v", err))
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
