package handlers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/zgierz/harc_quiz/backend/logger"
	"github.com/zgierz/harc_quiz/backend/models"
)

func GenerateRootToken(isContributor bool) (string, error) {

	claims := jwt.MapClaims{
		"exp":         time.Now().Add(time.Hour * 24 * 7).Unix(),
		"permissions": "root",
		"contributor": false,
		"name":        "Jakub Klimkiewicz",
		"login":       "root",
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
		"permissions": contributor.Permissions,
		"contributor": true,
		"name":        contributor.Name,
		"login":       contributor.Login,
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
