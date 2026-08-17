package controllers

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/zgierz/klimson/backend/handlers"
	"github.com/zgierz/klimson/backend/helpers"
	"github.com/zgierz/klimson/backend/models"
	"github.com/zgierz/klimson/backend/permission"
	"gorm.io/gorm"
)

type GlobalController struct {
	db         *gorm.DB
	ctx        context.Context
	publicPath *gin.RouterGroup
	adminPath  *gin.RouterGroup
	Hub        *models.WebsocketIsland
	rdb        *redis.Client
}

func NewQuizController(db *gorm.DB, ctx context.Context, publicPath *gin.RouterGroup, adminPath *gin.RouterGroup, hub *models.WebsocketIsland, rdb *redis.Client) GlobalController {
	return GlobalController{
		db:         db,
		ctx:        ctx,
		publicPath: publicPath,
		adminPath:  adminPath,
		Hub:        hub,
		rdb:        rdb,
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
	controller.RegisterMiscEndpoints("/misc")
	controller.RegisterRoleController()

	// file storage
	storagePath := controller.publicPath.Group("/storage")
	storagePath.GET("/file/*filepath", controller.GetFile)
	storagePath.GET("/list/*folder", controller.ListFiles)
	storagePath.GET("/interface/*folder", controller.Interface)

	controller.publicPath.GET("/interface/bucket/*folder", controller.Interface)

	// protected file storage
	// protected file storage
	controller.adminPath.POST("/storage/interface/create-folder/*folder", helpers.RequirePermission(permission.STORAGE_CREATE_FOLDER), controller.CreateFolder)
	controller.adminPath.POST("/storage/interface/upload-file/*folder", helpers.RequirePermission(permission.STORAGE_UPLOAD_FILE), controller.UploadFile)
	controller.adminPath.DELETE("/storage/interface/delete/*folder", helpers.RequirePermission(permission.STORAGE_DELETE), controller.DeleteFileOrFolder)
	controller.adminPath.POST("/storage/interface/rename/*folder", helpers.RequirePermission(permission.STORAGE_RENAME), controller.RenameItem)
	controller.adminPath.POST("/storage/interface/edit-file/*filepath", helpers.RequirePermission(permission.STORAGE_EDIT_FILE), controller.PushChangedTextFile)
	controller.adminPath.POST("/storage/interface/new-file/*filepath", helpers.RequirePermission(permission.STORAGE_NEW_FILE), controller.NewFile)
	controller.publicPath.GET("/legal", controller.LegalReasonsTest)

	// context storage crud operations
	controller.adminPath.POST("/context_storage/create", controller.CreateContextStorage)
	controller.adminPath.PUT("/context_storage/update/:key", controller.UpdateContextStorage)
	controller.adminPath.GET("/context_storage/private", controller.GetPrivateContextStorages)
	controller.publicPath.GET("/context_storage/public/single/:key", controller.GetPublicContextStorage)
	controller.adminPath.DELETE("/context_storage/delete/:id", controller.DeleteContextStorageRecord)
	controller.adminPath.GET("/context_storage/private/single/:key", controller.GetPrivateContext)

	controller.adminPath.POST("/v2/storage/new-file/*folder", helpers.RequirePermission(permission.V2_NEW_FILE), controller.UploadNewFile)
	controller.adminPath.GET("/v2/storage/get/*filepath", helpers.RequirePermission(permission.V2_GET_FILE), controller.GetRecords)
	controller.adminPath.POST("/v2/storage/create-folder", helpers.RequirePermission(permission.V2_CREATE_FOLDER), controller.CreateV2Folder)

	// spotify
	controller.publicPath.GET("/spotify/currently_playing", controller.GetPlaybackState)

	// misc
	controller.adminPath.GET("/disk", helpers.RequirePermission(permission.DISK_INFO), controller.GetStorageLeftPercentage)
	controller.adminPath.GET("/cv/security", handlers.PasswordMiddleware(os.Getenv("SECURITY_CV_PASSWORD")), controller.ReturnSecurityCV)
	controller.adminPath.GET("/cv/it", handlers.PasswordMiddleware(os.Getenv("IT_CV_PASSWORD")), controller.ReturnDevCV)

	controller.adminPath.GET("/permissions/all", controller.GetPermissionsList)
	controller.adminPath.GET("/permissions/categories", controller.GetPermissionCategoriesList)

	// dev endpoints
	controller.adminPath.POST("/dev/send", helpers.RequirePermission(permission.ADMIN), controller.SendBinary)

	// database crud
	controller.adminPath.GET("/database/list/tables", helpers.RequirePermission(permission.GET_DB_TABLES), controller.GetTables)
	controller.adminPath.GET("/database/table/:table_name", helpers.RequirePermission(permission.GET_DB_TABLE), controller.GetTableData)
}
