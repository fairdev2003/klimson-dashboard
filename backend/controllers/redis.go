package controllers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/zgierz/klimson/backend/api"
	"github.com/zgierz/klimson/backend/helpers"
	"github.com/zgierz/klimson/backend/logger"
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

func (controller GlobalController) RegisterRedisEndpoints(groupPrefix string) {

	redisGroupAdmin := controller.adminPath.Group(groupPrefix)
	redisGroupPublic := controller.publicPath.Group(groupPrefix)

	redisGroupAdmin.PUT("/set", helpers.RequirePermission(permission.SET_REDIS_KEY), controller.RDBSetKey)
	redisGroupAdmin.DELETE("del", helpers.RequirePermission(permission.DEL_REDIS_KEY), controller.RDBSetKey)
	redisGroupAdmin.GET("/key/info", helpers.RequirePermission(permission.REDIS_KEY_INFO), controller.RDBGetKeyInfo)

	redisGroupPublic.GET("/keys", controller.RDBGetAllExistingKeys)
	redisGroupPublic.GET("/get", controller.RDBGetKey)
	redisGroupPublic.GET("/ping", controller.PingRedis)

}
