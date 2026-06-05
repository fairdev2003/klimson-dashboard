package controllers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/zgierz/harc_quiz/backend/helpers"
	"github.com/zgierz/harc_quiz/backend/models"
	"gorm.io/gorm"
)

type GlobalController struct {
	db         *gorm.DB
	ctx        context.Context
	publicPath *gin.RouterGroup
	adminPath  *gin.RouterGroup
	Hub        *helpers.WSHub
	Files      []models.ListRecord
}

func NewQuizController(db *gorm.DB, ctx context.Context, publicPath *gin.RouterGroup, adminPath *gin.RouterGroup, hub *helpers.WSHub, files []models.ListRecord) GlobalController {
	return GlobalController{
		db:         db,
		ctx:        ctx,
		publicPath: publicPath,
		adminPath:  adminPath,
		Hub:        hub,
		Files:      files,
	}
}

func (controller GlobalController) RefreshCPU(ctx *gin.Context) {

	conn, err := helpers.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}

	controller.Hub.Mu.Lock()
	controller.Hub.Clients[conn] = true
	controller.Hub.Mu.Unlock()

	defer func() {
		controller.Hub.Mu.Lock()
		delete(controller.Hub.Clients, conn)
		controller.Hub.Mu.Unlock()
		conn.Close()
	}()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}
}

func (controller GlobalController) RegisterRoutes() {

	controller.publicPath.GET("/ws/stats/cpu", controller.RefreshCPU)

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
	controller.adminPath.POST("/storage/interface/rename/*folder", controller.RenameItem)
	controller.adminPath.GET("/storage/latest", controller.GetLatesFiles)

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

	// misc
	controller.adminPath.GET("/disk", controller.GetStorageLeftPercentage)

	// database crud
	controller.adminPath.GET("/database/list/tables", controller.GetTables)
	controller.adminPath.GET("/database/table/:table_name", controller.GetTableData)
}
