package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/zgierz/klimson/backend/config"
	"github.com/zgierz/klimson/backend/helpers"
	"github.com/zgierz/klimson/backend/logger"
	"github.com/zgierz/klimson/backend/models"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	server          *gin.Engine
	cfg             *models.Config
	db              *gorm.DB
	auth_token      string
	hashed_password string
	db_conn_string  string
	rdb             *redis.Client
)

func LoadConfig() {
	mainFolderEntries, err := os.ReadDir("./")
	if err != nil {
		fmt.Print(err.Error())
	}

	found := false
	for _, entry := range mainFolderEntries {
		found = false
		if entry.Name() == "server-config" {
			found = true
			break
		}

	}

	if !found {
		config, config_err := json.MarshalIndent(config.PreGeneratedConfig, "", "  ")
		if config_err != nil {
			fmt.Println("Error marshaling config:", err)
			return
		}

		os.Mkdir(serverConfigPath, 0755)
		err := os.WriteFile(serverConfigPath+"/config.json", config, 0644)
		if err != nil {
			fmt.Print(err.Error())
		}
		fmt.Print("\n")
		color.Green("✓ Config Created.")
	}

	cfg, err = config.GetConfig(mainPath, serverConfigPath)
	if err != nil {
		logger.ErrorLog(err.Error())
	}

	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	logger.GreenServerLog("✓ " + "Config successfully loaded!")
}

func Db() *gorm.DB {

	dsn := db_conn_string

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.ErrorLog(err.Error())
	}
	logger.GreenServerLog("✓ Sucessfully connected to the database")

	return DB
}

func Redis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("RDB_URL"), // Upewnij się, że tu jest "pro01.mikr.us:44285"
		Password: os.Getenv("RDB_PASS"),
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := client.Ping(ctx).Result()
	if err != nil {
		logger.ErrorLog("Error during connecting to redis database: %v", err)
	}

	logger.GreenServerLog("✓ Connected to redis storage successfully with message: ", response)
	return client
}

func AutoMigrateModels() {
	var modelsToMigrate []any
	for _, m := range models.MigratableModels {
		modelsToMigrate = append(modelsToMigrate, m.Model)
	}

	if err := db.AutoMigrate(modelsToMigrate...); err != nil {
		logger.ErrorLog("Error while migrating models to postgres database")
		return
	}
	logger.GreenServerLog("✓ " + "Migrated database models")
}

func StartGinServer() {
	server = gin.New()

	server.Static("/static", "/var/file_storage")
	InitRoutes()

	server.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": 404, "error": "Page not found", "message": "Page not found"})
	})

	logger.GreenServerLog("✓ " + cfg.DisplayName + " started listening on port " + strconv.Itoa(cfg.Port))
	if err := server.Run(":" + strconv.Itoa(cfg.Port)); err != nil {
		log.Fatalf("Failed to run server: %s", err)
	}
}

func FetchEnvVariables() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Nie udało się załadować pliku .env: %v", err)
	}

	hashed_password = os.Getenv("PASSWORD")
	auth_token = os.Getenv("TOKEN")
	db_conn_string = os.Getenv("DB_STRING")

	if hashed_password == "" || auth_token == "" || db_conn_string == "" {
		log.Fatal("Brakuje wymaganych zmiennych środowiskowych (PASSWORD lub TOKEN)")
	}

	logger.GreenServerLog("✓ Enviroment variables loaded correctly!")
}

func Init() {
	helpers.ClearConsole()

	LoadConfig()
	FetchEnvVariables()
	db = Db()
	rdb = Redis()
	AutoMigrateModels()
	StartGinServer()
}
