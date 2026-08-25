package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/zgierz/klimson/backend/controllers"
	"github.com/zgierz/klimson/backend/handlers"
	"github.com/zgierz/klimson/backend/helpers"
	"github.com/zgierz/klimson/backend/khttp"
	"github.com/zgierz/klimson/backend/logger"
	"github.com/zgierz/klimson/backend/models"

	"golang.org/x/crypto/bcrypt"
)

var (
	apiPath   *gin.RouterGroup
	wsPath    *gin.RouterGroup
	adminPath *gin.RouterGroup
)

func setToken(ctx *gin.Context, tokenName string, token string) {
	ctx.SetCookie(tokenName, token, 86400, "/", ".klimson.dev", false, true)
	ctx.SetCookie(tokenName, token, 86400, "/", "", false, true)
}

func InitRoutes() {

	ctx := context.Background()
	var origins []string = []string{"http://localhost:5173", "https://dashboard.klimson.dev", "https://klimson.dev", "http://mojprojekt.test:5173"}

	server.Use(helpers.NetworkLogger())
	server.Use(helpers.CorsConf(origins))

	apiPath = server.Group("/")
	wsPath = server.Group("/ws")
	adminPath = apiPath.Group("/admin")

	adminPath.Use(helpers.CorsConf(origins))
	adminPath.Use(AuthMiddleware())
	apiPath.Use(helpers.PublicRateLimiter())

	apiPath.GET("/", func(c *gin.Context) {
		c.JSON(200, "Nie powinno ciebie tu byc ://")
	})
	apiPath.GET("/callback", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"succes": true,
		})
	})

	type MostCommonColorBody struct {
		ImageURL string `json:"image_url" binding:"required"`
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
	cpu_hub := helpers.NewHub("cpu")
	logger_ws := helpers.NewHub("logger")
	go cpu_hub.Run()
	go logger_ws.Run()

	go func() {
		ticker := time.NewTicker(2 * time.Second)
		for range ticker.C {
			// Pobieranie użycia CPU
			cpuUsage, _ := cpu.Percent(0, false)

			// Pobieranie statystyk pamięci RAM
			memInfo, err := mem.VirtualMemory()
			var memUsage float64 = 0
			var totalMem uint64 = 0
			var usedMem uint64 = 0

			if err == nil {
				memUsage = memInfo.UsedPercent
				totalMem = memInfo.Total
				usedMem = memInfo.Used
			}

			// Wysyłanie danych do hubu
			cpu_hub.Send(map[string]interface{}{
				"cpu":       cpuUsage,
				"memory":    memUsage,
				"mem_total": totalMem,
				"mem_used":  usedMem,
			})
		}
	}()
	controller := controllers.NewQuizController(db, ctx, apiPath, adminPath, &models.WebsocketIsland{
		CPUHub:    (*models.WSHub)(cpu_hub),
		LoggerHub: (*models.WSHub)(logger_ws),
	}, rdb)
	controller.RegisterRoutes()

	apiPath.POST("/login", func(c *gin.Context) {
		var req LoginRequest
		var user models.Contributor

		if err := c.ShouldBindJSON(&req); err != nil {
			logger.ErrorLog("Niepoprawny format JSONA")
			khttp.BadRequestResponse(c, gin.H{}, "Invalid form data")
			return
		}

		result := db.Model(models.Contributor{}).Where("login = ?", req.Login).Find(&user)

		if result.RowsAffected == 0 {
			err := bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(req.Password))
			if err != nil {
				logger.ErrorLog("Incorrect password")
				khttp.UnauthorizedResponse(c, gin.H{}, "Invalid password")
				return
			}

			token, err := handlers.GenerateRootToken(false)

			setToken(c, "k-token", token)

			c.Writer.Header().Set("X-Token", token)

			if err != nil {
				logger.ErrorLog("Error while generating token:", err.Error())
				khttp.InternalServerErrorResponse(c, nil, "Error while generating token:", err.Error())
			}

			config, err := controllers.GlobalController.GetUserConfigFromRdbHash(controller, "root")

			khttp.SuccessResponse(c, gin.H{"user_config": config}, "Login was successfull")

		} else {
			logger.ServerLog("Znaleziono kontrybutora o nicku: ", user.Password)
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
			if err != nil {
				logger.ErrorLog("Incorrect password")
				khttp.BadRequestResponse(c, gin.H{}, "Invalid password")
				return
			}

			token, err := handlers.GenerateToken(user)
			if err != nil {
				logger.ErrorLog("Error while generating token:", err.Error())
				khttp.InternalServerErrorResponse(c, nil, "Error while generating token:", err.Error())
				return
			}

			c.SetCookie("k-token", token, 3600, "/", "", false, true)

			khttp.SuccessResponse(c, gin.H{}, "Login was successfull")
		}

	})

	apiPath.POST("/auth/logout", func(ctx *gin.Context) {
		ctx.SetCookie("k-token", "", -1, "/", "", false, true)

		khttp.SuccessResponse(ctx, gin.H{"success": true}, "Cookie is successfully deleted!")
	})

}
