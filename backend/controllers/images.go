package controllers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (gc GlobalController) SendImage(ctx *gin.Context) {
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savingKey := ctx.Query("key")
	newID := uuid.New().String()
	ext := filepath.Ext(file.Filename)
	newFileName := newID + ext

	savePath := "static/uploads/" + savingKey + "/" + newFileName
	err = ctx.SaveUploadedFile(file, savePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":   "Upload successful",
		"file_name": newFileName,
		"id":        newID,
	})
}

func (gc GlobalController) ListImages(ctx *gin.Context) {
	key := ctx.Query("key")
	dir := "static/uploads/" + key

	files, err := os.ReadDir(dir)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var fileList []string
	for _, f := range files {
		if !f.IsDir() {
			fileList = append(fileList, f.Name())
		}
	}

	ctx.JSON(http.StatusOK, fileList)
}

func (gc GlobalController) DeleteImage(ctx *gin.Context) {
	key := ctx.Query("key")
	fileName := ctx.Query("file")

	if key == "" || fileName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing key or file parameter"})
		return
	}

	filePath := "static/uploads/" + key + "/" + fileName

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	if err := os.Remove(filePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete file"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "File deleted successfully",
		"file":    fileName,
	})
}
