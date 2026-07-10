package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/zgierz/klimson/backend/khttp"
	"github.com/zgierz/klimson/backend/logger"
	"github.com/zgierz/klimson/backend/models"
)

func (gc GlobalController) UploadNewFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Brak pliku w żądaniu", "success": false})
		return
	}

	folderPath := ctx.Param("folder")
	dst := filepath.Join("./static/private", folderPath, file.Filename)

	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "success": false})
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
		fmt.Printf("BŁĄD BAZY: %v\n", errDB) // Sprawdź konsolę serwera!
		khttp.InternalServerErrorResponse(ctx, nil, errDB.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Plik przesłany pomyślnie", "success": true})
}

// erro := gc.db.Create(&models.Storage{
// 		Name:        meta.Name,
// 		Description: meta.Description,
// 		Filename:    header.Filename,
// 		Path:        path,
// 		Filetype:    header.Header.Get("Content-Type"),
// 	}).Error
// 	if erro != nil {
// 		khttp.InternalServerErrorResponse(ctx, nil, erro.Error())
// 		return
// 	}

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
		khttp.InternalServerErrorResponse(ctx, nil, err.Error())
	}

	ctx.JSON(200, storage)
}

type CreateFolderRequest struct {
	Name string `json:"name" binding:"required"` // 'required' wymusza podanie tej wartości
	Path string `json:"path" binding:"required"`
}

func (gc GlobalController) CreateV2Folder(ctx *gin.Context) {
	var req CreateFolderRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Niepoprawne dane wejściowe: " + err.Error(),
		})
		return
	}

	// Teraz możesz bezpiecznie używać zmiennych req.Name oraz req.Path
	fullPath := filepath.Join("./static/private", req.Path, req.Name)

	err := os.MkdirAll(fullPath, 0755)
	if err != nil {
		khttp.InternalServerErrorResponse(ctx, nil, "Nie udało się stworzyć folderu: "+err.Error())
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
		khttp.InternalServerErrorResponse(ctx, nil, "Błąd zapisu w bazie: "+errDB.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Folder " + req.Name + " utworzony"})
}

func (gc GlobalController) RegisterV2StorageEndpoints() {

}
