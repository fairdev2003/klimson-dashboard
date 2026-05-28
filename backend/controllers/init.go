package controllers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/zgierz/harc_quiz/backend/helpers"
	"gorm.io/gorm"
)

type GlobalController struct {
	db         *gorm.DB
	ctx        context.Context
	publicPath *gin.RouterGroup
	adminPath  *gin.RouterGroup
	Hub        *helpers.WSHub
}

func NewQuizController(db *gorm.DB, ctx context.Context, publicPath *gin.RouterGroup, adminPath *gin.RouterGroup, hub *helpers.WSHub) GlobalController {
	return GlobalController{
		db:         db,
		ctx:        ctx,
		publicPath: publicPath,
		adminPath:  adminPath,
		Hub:        hub,
	}
}

func (controller GlobalController) RegisterRoutes() {

	// file storage
	storagePath := controller.publicPath.Group("/storage")
	storagePath.GET("/file/*filepath", controller.GetFile)
	storagePath.GET("/list/*folder", controller.ListFiles)
	storagePath.GET("/interface/*folder", controller.Interface)
	controller.publicPath.GET("/interface/bucket/*folder", controller.Interface)

	// protected file storage
	controller.adminPath.POST("/storage/interface/create-folder/*folder", controller.CreateFolder)
	controller.adminPath.POST("/storage/interface/upload-file/*folder", controller.UploadFile)
	controller.adminPath.DELETE("/storage/interface/delete/*folder", controller.DeleteFileOrFolder)

	// context storage crud operations
	controller.adminPath.POST("/context_storage/create", controller.CreateContextStorage)
	controller.adminPath.PUT("/context_storage/update/:key", controller.UpdateContextStorage)
	controller.adminPath.GET("/context_storage/private", controller.GetPrivateContextStorages)
	controller.publicPath.GET("/context_storage/public/single/:key", controller.GetPublicContextStorage)
	controller.adminPath.DELETE("/context_storage/delete/:id", controller.DeleteContextStorageRecord)
	controller.adminPath.GET("/context_storage/private/single/:key", controller.GetPrivateContext)

	// pg3d
	controller.publicPath.GET("/pg3d/clan_info/:clan_id", controller.GetClanInfo)
	controller.publicPath.GET("/pg3d/player_data/:player_id", controller.GetPlayerData)

	// spotify
	controller.publicPath.GET("/spotify/currently_playing", controller.GetPlaybackState)

	// database crud
	controller.adminPath.GET("/database/list/tables", controller.GetTables)
	controller.adminPath.GET("/database/table/:table_name", controller.GetTableData)
}
