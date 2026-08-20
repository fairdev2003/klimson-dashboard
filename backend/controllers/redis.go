package controllers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/zgierz/klimson/backend/api"
	"github.com/zgierz/klimson/backend/helpers"
	"github.com/zgierz/klimson/backend/logger"
	"github.com/zgierz/klimson/backend/models"
	"github.com/zgierz/klimson/backend/permission"
)

func (controller GlobalController) PingRedis(ctx *gin.Context) {

	response, err := controller.rdb.Ping(controller.ctx).Result()

	if err != nil {
		api.InternalServerErrorResponse(ctx, err)
	}

	api.SuccessResponse(ctx, gin.H{"ping": response})
}

func (controller GlobalController) RDBSetKey(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	logger.ServerLog("Header: ", header)
	key := ctx.Query("key")
	value := ctx.Query("value")

	result, err := controller.rdb.Set(controller.ctx, key, value, 0).Result()
	if err != nil {
		api.InternalServerErrorResponse(ctx, nil, err.Error())
		return
	}

	api.SuccessResponse(ctx, gin.H{"result": result}, "Success!")
}

func (controller GlobalController) RDBGetKey(ctx *gin.Context) {
	key := ctx.Query("key")

	result, err := controller.rdb.Get(ctx, key).Result()
	if err != nil {
		api.InternalServerErrorResponse(ctx, nil, err.Error())
		return
	}

	api.SuccessResponse(ctx, gin.H{"result": result}, "Success!")
}

func (controller GlobalController) RDBDelKey(ctx *gin.Context) {
	key := ctx.Query("key")

	result, err := controller.rdb.Del(ctx, key).Result()
	if err != nil {
		api.InternalServerErrorResponse(ctx, nil, err.Error())
		return
	}

	api.SuccessResponse(ctx, gin.H{"result": result}, "Success!")
}

func getAllKeys(ctx context.Context, client *redis.Client) ([]string, error) {
	var cursor uint64
	var allKeys []string

	for {
		keys, nextCursor, err := client.Scan(ctx, cursor, "*", 10).Result()
		if err != nil {
			return nil, err
		}

		allKeys = append(allKeys, keys...)
		cursor = nextCursor

		if cursor == 0 {
			break
		}
	}

	return allKeys, nil
}

func (controller GlobalController) RDBGetAllExistingKeys(ctx *gin.Context) {
	keys, err := getAllKeys(controller.ctx, controller.rdb)

	if err != nil {
		api.InternalServerErrorResponse(ctx, nil, err.Error())
		return
	}

	api.SuccessResponse(ctx, gin.H{"rdbs": keys}, "Success!")
}

func (controller GlobalController) RDBGetKeyInfo(ctx *gin.Context) {
	key := ctx.Query("key")

	ttl, _ := controller.rdb.TTL(ctx, key).Result()

	mem, _ := controller.rdb.MemoryUsage(ctx, key, 0).Result()

	kType, _ := controller.rdb.Type(ctx, key).Result()

	idle, _ := controller.rdb.ObjectIdleTime(ctx, key).Result()

	api.SuccessResponse(ctx, gin.H{"memory_usage": mem, "type": kType, "ttl": ttl, "idle": idle}, "Request was successfull")

}

func (controller GlobalController) GetUserConfigFromRdbHash(userID string) (*models.ClientConfig, error) {
	redisKey := fmt.Sprintf("user:config:%s", userID)

	resultMap, err := controller.rdb.HGetAll(controller.ctx, redisKey).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get client config from redis: %w", err)
	}

	if len(resultMap) == 0 {
		return nil, fmt.Errorf("config not found for user %s", userID)
	}

	config := &models.ClientConfig{}
	config.DashboardTheme = resultMap["theme"]
	config.CodeTheme = resultMap["code_theme"]
	config.SidebarBehavior = resultMap["sidebarBehavior"]

	if pillsStr, exists := resultMap["client_pills"]; exists && pillsStr != "" {
		_ = json.Unmarshal([]byte(pillsStr), &config.SidebarPreferences)
	}

	if dockStr, exists := resultMap["dock_on"]; exists {
		val := dockStr == "true" || dockStr == "1"
		config.Dock = &val
	}

	if bookmarksStr, exists := resultMap["bookmarks"]; exists && bookmarksStr != "" {
		_ = json.Unmarshal([]byte(bookmarksStr), &config.Bookmarks)
	}

	return config, nil
}

func (controller GlobalController) RDBGetUserConfig(ctx *gin.Context) {
	user_id := ctx.Query("user_id")

	if user_id == "" {
		api.BadRequestResponse(ctx, nil, "Invalid request. Query param 'user_id' is required.")
		return
	}

	config, err := controller.GetUserConfigFromRdbHash(user_id)

	if err != nil {
		api.InternalServerErrorResponse(ctx, nil, fmt.Sprint("Error occured whil retrieving config from redis database: %s", err.Error()))
		return
	}

	api.SuccessResponse(ctx, gin.H{"client_config": config})
}

func (controller GlobalController) saveClientConfig(userID string, config models.ClientConfig) error {
	redisKey := fmt.Sprintf("user:config:%s", userID)

	dockOnVal := false
	if config.Dock != nil {
		dockOnVal = *config.Dock
	}

	sidebarJSON, err := json.Marshal(config.SidebarPreferences)
	if err != nil {
		return fmt.Errorf("failed to marshal sidebar preferences: %w", err)
	}

	bookmarksJSON, err := json.Marshal(config.Bookmarks)
	if err != nil {
		return fmt.Errorf("failed to marshal bookmarks: %w", err)
	}

	values := map[string]interface{}{
		"theme":           config.DashboardTheme,
		"code_theme":      config.CodeTheme,
		"client_pills":    string(sidebarJSON),
		"dock_on":         dockOnVal,
		"bookmarks":       string(bookmarksJSON),
		"sidebarBehavior": config.SidebarBehavior,
	}

	err = controller.rdb.HSet(controller.ctx, redisKey, values).Err()
	if err != nil {
		return fmt.Errorf("failed to save client config to redis hash: %w", err)
	}

	return nil
}

func (controller GlobalController) RDBSetUserConfig(ctx *gin.Context) {
	user_id := ctx.Query("user_id")

	if user_id == "" {
		api.BadRequestResponse(ctx, nil, "Invalid request. Query param 'user_id' is required.")
		return
	}

	var config models.ClientConfig

	if err := ctx.ShouldBindJSON(&config); err != nil {
		api.BadRequestResponse(ctx, nil, fmt.Sprintf("Invalid request body: %s", err.Error()))
		return
	}

	err := controller.saveClientConfig(user_id, config)
	if err != nil {
		api.InternalServerErrorResponse(ctx, nil, fmt.Sprintf("Error occurred while saving config to redis database: %s", err.Error()))
		return
	}

	api.SuccessResponse(ctx, gin.H{"client_config": config}, "User config successfully updated")
}

func (controller GlobalController) RegisterRedisEndpoints(groupPrefix string) {

	redisGroupAdmin := controller.adminPath.Group(groupPrefix)
	redisGroupPublic := controller.publicPath.Group(groupPrefix)

	redisGroupAdmin.PUT("/set", helpers.RequirePermission(permission.SET_REDIS_KEY), controller.RDBSetKey)
	redisGroupAdmin.DELETE("del", helpers.RequirePermission(permission.DEL_REDIS_KEY), controller.RDBSetKey)
	redisGroupAdmin.GET("/key/info", helpers.RequirePermission(permission.REDIS_KEY_INFO), controller.RDBGetKeyInfo)

	// /admin/user-config/get?user_id=1
	redisGroupAdmin.GET("/user-config/get", controller.RDBGetUserConfig)
	redisGroupAdmin.POST("/user-config/set", controller.RDBSetUserConfig)

	redisGroupPublic.GET("/keys", controller.RDBGetAllExistingKeys)
	redisGroupPublic.GET("/get", controller.RDBGetKey)
	redisGroupPublic.GET("/ping", controller.PingRedis)

}
