package controllers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/zgierz/klimson/backend/helpers"
	"github.com/zgierz/klimson/backend/models"
	"gorm.io/gorm"
)

type GlobalController struct {
	db         *gorm.DB
	ctx        context.Context
	publicPath *gin.RouterGroup
	adminPath  *gin.RouterGroup
	Hub        *models.WebsocketIsland
	Files      []models.ListRecord
	rdb        *redis.Client
	state      *helpers.StateHub
}

func NewQuizController(db *gorm.DB, ctx context.Context, publicPath *gin.RouterGroup, adminPath *gin.RouterGroup, hub *models.WebsocketIsland, files []models.ListRecord, rdb *redis.Client, state *helpers.StateHub) GlobalController {
	return GlobalController{
		db:         db,
		ctx:        ctx,
		publicPath: publicPath,
		adminPath:  adminPath,
		Hub:        hub,
		Files:      files,
		rdb:        rdb,
		state:      state,
	}
}

func (controller GlobalController) RefreshCPU(ctx *gin.Context) {

	conn, err := helpers.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}

	controller.Hub.CPUHub.Mu.Lock()
	controller.Hub.CPUHub.Clients[conn] = true
	controller.Hub.CPUHub.Mu.Unlock()

	defer func() {
		controller.Hub.CPUHub.Mu.Lock()
		delete(controller.Hub.CPUHub.Clients, conn)
		controller.Hub.CPUHub.Mu.Unlock()
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

	controller.RegisterV2StorageEndpoints()
	controller.RegisterUserController()
	controller.RegisterRedisEndpoints("/redis")
	controller.RegisterStateEndpoints("/state")

	// file storage
	storagePath := controller.publicPath.Group("/storage")
	storagePath.GET("/file/*filepath", controller.GetFile)
	storagePath.GET("/list/*folder", controller.ListFiles)
	storagePath.GET("/interface/*folder", controller.Interface)

	controller.publicPath.GET("/interface/bucket/*folder", controller.Interface)

	// protected file storage
	controller.adminPath.POST("/storage/interface/create-folder/*folder", helpers.RequirePermission(
		helpers.PermissionsMetadata{
			Name:          "Create Private Folder",
			Icon:          "chuj go wie",
			PermissionTag: "storage:private:create-folder",
			Color:         "bg-blue-500",
			Description:   "",
		},
	), controller.CreateFolder)
	controller.adminPath.POST("/storage/interface/upload-file/*folder", controller.UploadFile)
	controller.adminPath.DELETE("/storage/interface/delete/*folder", controller.DeleteFileOrFolder)
	controller.adminPath.POST("/storage/interface/rename/*folder", controller.RenameItem)
	controller.adminPath.POST("/storage/interface/edit-file/*filepath", controller.PushChangedTextFile)
	controller.adminPath.POST("/storage/interface/new-file/*filepath", controller.NewFile)
	controller.adminPath.GET("/storage/latest", helpers.RequirePermission(
		helpers.PermissionsMetadata{
			Name:          "Latest Files",
			Icon:          "chuj go wie",
			PermissionTag: "storage:latest-files",
			Color:         "bg-blue-500",
			Description:   "",
		},
	), controller.GetLatesFiles)

	controller.publicPath.GET("/legal", controller.LegalReasonsTest)

	// context storage crud operations
	controller.adminPath.POST("/context_storage/create", controller.CreateContextStorage)
	controller.adminPath.PUT("/context_storage/update/:key", controller.UpdateContextStorage)
	controller.adminPath.GET("/context_storage/private", controller.GetPrivateContextStorages)
	controller.publicPath.GET("/context_storage/public/single/:key", controller.GetPublicContextStorage)
	controller.adminPath.DELETE("/context_storage/delete/:id", controller.DeleteContextStorageRecord)
	controller.adminPath.GET("/context_storage/private/single/:key", controller.GetPrivateContext)

	controller.adminPath.POST("/v2/storage/new-file/*folder", controller.UploadNewFile)
	controller.adminPath.GET("/v2/storage/get/*filepath", controller.GetRecords)
	controller.adminPath.POST("/v2/storage/create-folder", controller.CreateV2Folder)

	// pg3d
	controller.publicPath.GET("/pg3d/clan_info/:clan_id", controller.GetClanInfo)
	controller.publicPath.GET("/pg3d/player_data/:player_id", controller.GetPlayerData)
	controller.publicPath.GET("/pg3d/clan/valor_history/:clan_id", controller.GetValorHistory)

	// spotify
	controller.publicPath.GET("/spotify/currently_playing", controller.GetPlaybackState)

	// misc
	controller.adminPath.GET("/disk", helpers.RequirePermission(
		helpers.PermissionsMetadata{
			Name:          "Disk Info",
			Icon:          "vaadin:harddrive",
			PermissionTag: "disk:info",
			Color:         "bg-blue-500",
			Description:   "",
		},
	), controller.GetStorageLeftPercentage)

	controller.adminPath.GET("/permissions/all", controller.GetPermissionsList)

	// database crud
	controller.adminPath.GET("/database/list/tables", controller.GetTables)
	controller.adminPath.GET("/database/table/:table_name", controller.GetTableData)
}
