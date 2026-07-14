package main

import (
	"context"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/zgierz/klimson/backend/controllers"
	"github.com/zgierz/klimson/backend/handlers"
	"github.com/zgierz/klimson/backend/helpers"
	"github.com/zgierz/klimson/backend/logger"
	"github.com/zgierz/klimson/backend/models"
	"golang.org/x/crypto/bcrypt"
)

var (
	apiPath     *gin.RouterGroup
	wsPath      *gin.RouterGroup
	adminPath   *gin.RouterGroup
	latestFiles []models.ListRecord
)

func AddRandomRecords() {
	names := []string{"dokument.txt", "zdjecie.jpg", "projekt.go", "backup.zip", "notatki.md"}

	for i := 0; i < 3; i++ {
		record := models.ListRecord{
			Name:         names[rand.Intn(len(names))],
			IsDir:        rand.Intn(2) == 0,
			ModifiedTime: time.Now().Add(time.Duration(-rand.Intn(1000)) * time.Hour),
			Size:         rand.Int63n(1024 * 1024),
		}
		latestFiles = append(latestFiles, record)
	}
}

func InitRoutes() {

	ctx := context.Background()
	var origins []string = []string{"http://localhost:5173", "https://dashboard.klimson.dev"}
	server.Use(helpers.CorsConf(origins))
	server.Use(helpers.NetworkLogger())
	apiPath = server.Group("/")
	wsPath = server.Group("/ws")
	adminPath = apiPath.Group("/admin")
	adminPath.Use(AuthMiddleware())
	adminPath.Use(helpers.CorsConf(origins))

	// root routes
	apiPath.GET("/", func(c *gin.Context) {
		c.JSON(200, "Nie powinno ciebie tu byc ://")
	})
	apiPath.GET("/callback", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"succes": true,
		})
	})

	type MostCommonColorBody struct {
		ImageURL string `json:"image_url" binding:"required"` // Dodaliśmy tag 'required' dla walidacji
	}

	apiPath.POST("/most_common_image_color", func(ctx *gin.Context) {
		var body MostCommonColorBody

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Niepoprawny format JSON lub brak pola image_url",
			})
			return
		}

		hex, err := helpers.GetMostCommonColor(body.ImageURL)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"hex": hex,
		})
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

			c.SetCookie("token", token, 3600, "/", "", false, true)
			c.Writer.Header().Set("X-Token", token)

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
				c.JSON(500, gin.H{"text": "błąd serwera przy generowaniu tokena"})
				return
			}

			c.SetCookie("token", token, 3600, "/", "", false, true)
			c.Writer.Header().Set("X-Token", token)

			c.JSON(200, gin.H{"token": ""})
		}

	})

	cpu_hub := helpers.NewHub("cpu")
	logger_ws := helpers.NewHub("logger")
	go cpu_hub.Run()
	go logger_ws.Run()

	go func() {
		ticker := time.NewTicker(2 * time.Second)
		for range ticker.C {
			cpuUsage, _ := cpu.Percent(0, false)

			cpu_hub.Send(map[string]float64{"cpu": cpuUsage[0]})
		}
	}()
	AddRandomRecords()
	newQuizController := controllers.NewQuizController(db, ctx, apiPath, adminPath, cpu_hub, latestFiles, rdb)
	newQuizController.RegisterRoutes()
}
