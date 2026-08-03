package controllers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/zgierz/klimson/backend/api"
	"github.com/zgierz/klimson/backend/logger"
	"github.com/zgierz/klimson/backend/models"
)

func (gc GlobalController) UploadNewFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")

	if err != nil {
		api.BadRequestResponse(ctx, nil, "Brak pliku w żądaniu")
		return
	}

	folderPath := ctx.Param("folder")
	dst := filepath.Join("./static/private", folderPath, file.Filename)

	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		api.InternalServerErrorResponse(ctx, nil, err.Error())
		return
	}

	storageEntry := &models.Storage{
		Name:        file.Filename,
		Description: "",
		Filename:    file.Filename,
		Path:        folderPath,
		Filetype:    file.Header.Get("Content-Type"),
	}

	errDB := gc.db.Debug().Create(storageEntry).Error
	if errDB != nil {
		fmt.Printf("BŁĄD BAZY: %v\n", errDB)
		api.InternalServerErrorResponse(ctx, nil, errDB.Error())
		return
	}

	api.SuccessResponse(ctx, nil, "Plik przesłany pomyślnie")
}

func (gc GlobalController) GetV2File(c *gin.Context) {
	filePath := c.Param("filepath")

	fullPath := "/" + filePath

	c.File(fullPath)
}

func (gc GlobalController) GetRecords(ctx *gin.Context) {
	filePath := ctx.Param("filepath")
	logger.ServerLog(filePath)

	var storage []models.Storage

	err := gc.db.Where("path = ?", filePath).Find(&storage).Error
	if err != nil {
		api.InternalServerErrorResponse(ctx, nil, err.Error())
		return
	}

	api.SuccessResponse(ctx, storage)
}

type CreateFolderRequest struct {
	Name string `json:"name" binding:"required"`
	Path string `json:"path" binding:"required"`
}

func (gc GlobalController) CreateV2Folder(ctx *gin.Context) {
	var req CreateFolderRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		api.BadRequestResponse(ctx, nil, "Niepoprawne dane wejściowe: "+err.Error())
		return
	}

	fullPath := filepath.Join("./static/private", req.Path, req.Name)

	err := os.MkdirAll(fullPath, 0755)
	if err != nil {
		api.InternalServerErrorResponse(ctx, nil, "Nie udało się stworzyć folderu: "+err.Error())
		return
	}

	newFolder := &models.Storage{
		Name:        req.Name,
		Description: "Folder utworzony przez panel",
		Filename:    "",
		Filetype:    "folder",
		Path:        req.Path,
	}

	errDB := gc.db.Create(newFolder).Error
	if errDB != nil {
		api.InternalServerErrorResponse(ctx, nil, "Błąd zapisu w bazie: "+errDB.Error())
		return
	}

	api.SuccessResponse(ctx, nil, "Folder "+req.Name+" utworzony")
}

func (gc GlobalController) RegisterV2StorageEndpoints() {

}
