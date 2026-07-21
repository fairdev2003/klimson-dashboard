package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zgierz/klimson/backend/khttp"
	"github.com/zgierz/klimson/backend/logger"
)

func (controller GlobalController) GetState(ctx *gin.Context) {
	key := ctx.Param("key")

	state, ok := controller.state.Get(key)
	if ok {
		khttp.SuccessResponse(ctx, gin.H{"state": state}, "Key found!")
	} else {
		khttp.InternalServerErrorResponse(ctx, nil, "No key found in global state")
	}
}

func (controller GlobalController) SetState(ctx *gin.Context) {
	type PostData struct {
		Key   string      `json:"key"`
		Value interface{} `json:"value"`
	}

	var body PostData

	if err := ctx.ShouldBindJSON(&body); err != nil {
		khttp.BadRequestResponse(ctx, nil, "No key found in global state")
		return
	}

	controller.state.Set(body.Key, body.Value)
	stateValue, _ := controller.state.Get(body.Key)
	logger.GreenServerLog(stateValue)
	khttp.SuccessResponse(ctx, gin.H{"state": stateValue}, "State updated successfully")
}

func (controller GlobalController) RegisterStateEndpoints(groupPrefix string) {
	// stateGroupAdmin := controller.adminPath.Group(groupPrefix)
	stateGroupPublic := controller.publicPath.Group(groupPrefix)
	stateGroupPublic.GET("/get/:key", controller.GetState)
	stateGroupPublic.POST("/set", controller.SetState)
}
