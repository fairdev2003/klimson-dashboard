package controllers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/zgierz/klimson/backend/khttp"
)

func (controller GlobalController) PingRedis(ctx *gin.Context) {

	response, err := controller.rdb.Ping(controller.ctx).Result()

	if err != nil {
		khttp.InternalServerErrorResponse(ctx, err)
	}

	khttp.SuccessResponse(ctx, gin.H{"ping": response})
}

func (controller GlobalController) RDBSetKey(ctx *gin.Context) {
	key := ctx.Query("key")
	value := ctx.Query("value")

	result, err := controller.rdb.Set(controller.ctx, key, value, 0).Result()
	if err != nil {
		khttp.InternalServerErrorResponse(ctx, nil, err.Error())
		return
	}

	khttp.SuccessResponse(ctx, gin.H{"result": result}, "Success!")
}

func (controller GlobalController) RDBGetKey(ctx *gin.Context) {
	key := ctx.Query("key")

	result, err := controller.rdb.Get(ctx, key).Result()
	if err != nil {
		khttp.InternalServerErrorResponse(ctx, nil, err.Error())
		return
	}

	khttp.SuccessResponse(ctx, gin.H{"result": result}, "Success!")
}

func (controller GlobalController) RDBDelKey(ctx *gin.Context) {
	key := ctx.Query("key")

	result, err := controller.rdb.Del(ctx, key).Result()
	if err != nil {
		khttp.InternalServerErrorResponse(ctx, nil, err.Error())
		return
	}

	khttp.SuccessResponse(ctx, gin.H{"result": result}, "Success!")
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
		khttp.InternalServerErrorResponse(ctx, nil, err.Error())
		return
	}

	khttp.SuccessResponse(ctx, gin.H{"rdbs": keys}, "Success!")
}

func (controller GlobalController) RDBGetKeyInfo(ctx *gin.Context) {
	key := ctx.Query("key")

	ttl, _ := controller.rdb.TTL(ctx, key).Result()

	mem, _ := controller.rdb.MemoryUsage(ctx, key, 0).Result()

	kType, _ := controller.rdb.Type(ctx, key).Result()

	idle, _ := controller.rdb.ObjectIdleTime(ctx, key).Result()

	khttp.SuccessResponse(ctx, gin.H{"memory_usage": mem, "type": kType, "ttl": ttl, "idle": idle}, "Request was successfull")

}

func (controller GlobalController) StartRedisEndpoints() {
	controller.publicPath.GET("/redis/ping", controller.PingRedis)
	controller.publicPath.GET("/redis/get", controller.RDBGetKey)
	controller.adminPath.PUT("/redis/set", controller.RDBSetKey)
	controller.adminPath.DELETE("/redis/del", controller.RDBSetKey)
	controller.publicPath.GET("/redis/keys", controller.RDBGetAllExistingKeys)
	controller.publicPath.GET("/redis/key/info", controller.RDBGetKeyInfo)
}
