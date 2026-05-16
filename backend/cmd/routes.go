package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zgierz/harc_quiz/backend/controllers"
	"github.com/zgierz/harc_quiz/backend/handlers"
	"github.com/zgierz/harc_quiz/backend/helpers"
	"github.com/zgierz/harc_quiz/backend/logger"
	"github.com/zgierz/harc_quiz/backend/models"
	"golang.org/x/crypto/bcrypt"
)

var (
	apiPath   *gin.RouterGroup
	wsPath    *gin.RouterGroup
	adminPath *gin.RouterGroup
)

func InitRoutes() {

	ctx := context.Background()
	server.Use(helpers.CorsConf("*"))
	server.Use(helpers.NetworkLogger())
	apiPath = server.Group(cfg.Api.Version).Group("/api")
	wsPath = server.Group(cfg.Api.Version).Group("/ws")
	adminPath = apiPath.Group("/admin")
	adminPath.Use(AuthMiddleware())
	adminPath.Use(helpers.CorsConf("*"))

	// root routes
	apiPath.GET("/", func(c *gin.Context) {
		c.JSON(200, "Nie powinno ciebie tu byc ://")
	})
	adminPath.GET("/verify", func(c *gin.Context) {
		c.JSON(200, gin.H{"access": true})
	})
	wsPath.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"text": "WS Route"})
	})
	adminPath.GET("/routes", func(c *gin.Context) {
		type Route struct {
			Method string `json:"method"`
			Path   string `json:"path"`
		}
		var routesList []Route
		routes := server.Routes()
		for _, i := range routes {
			path := i.Path
			route := Route{Method: i.Method, Path: path}
			routesList = append(routesList, route)
		}
		c.JSON(200, routesList)
	})

	type LoginRequest struct {
		Password string `json:"password"`
		Login    string `json:"login"`
	}

	apiPath.POST("/login", func(c *gin.Context) {
		var req LoginRequest
		var user models.Contributor

		if err := c.ShouldBindJSON(&req); err != nil {
			logger.ErrorLog("Niepoprawny format JSONA")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Błędny JSON"})
			return
		}

		result := db.Model(models.Contributor{}).Where("login = ?", req.Login).Find(&user)

		if result.RowsAffected == 0 {
			err := bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(req.Password))
			if err != nil {
				logger.ErrorLog("Incorrect password")
				c.JSON(400, gin.H{"text": "niepoprawne hasło"})
				return
			}

			token, err := handlers.GenerateRootToken(false)
			if err != nil {
				logger.ErrorLog("Error while generating token:", err.Error())
			}

			c.JSON(200, gin.H{"token": token})
		} else {
			logger.ServerLog("Znaleziono kontrybutora o nicku: ", user.Password)
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
			if err != nil {
				logger.ErrorLog("Incorrect password")
				c.JSON(400, gin.H{"text": "niepoprawne hasło"})
				return
			}

			token, err := handlers.GenerateToken(user)
			if err != nil {
				logger.ErrorLog("Error while generating token:", err.Error())
			}

			c.JSON(200, gin.H{"token": token})
		}

	})

	hub := helpers.NewHub()
	go hub.Run()
	newQuizController := controllers.NewQuizController(db, ctx, apiPath, adminPath, hub)
	newQuizController.RegisterRoutes()
}
