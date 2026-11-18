package controllers

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zgierz/klimson/backend/api"
)

func (gc GlobalController) ServeFile(ctx *gin.Context) {
	// Pobieranie parametru ze ścieżki (np. /public-drive/download/km1.zip -> "km1.zip")
	filePath := ctx.Param("path")

	if filePath == "" {
		api.BadRequestResponse(ctx, nil, "No path defined")
		return
	}

	baseStorageDir := "/home/klimson/file_storage"
	fullPath := filepath.Join(baseStorageDir, filePath)
	cleanPath := filepath.Clean(fullPath)

	rel, err := filepath.Rel(baseStorageDir, cleanPath)
	if err != nil || strings.HasPrefix(rel, "..") {
		api.ForbiddenResponse(ctx, nil, "Access denied")
		return
	}

	info, err := os.Stat(cleanPath)
	if os.IsNotExist(err) || info.IsDir() {
		api.NotFoundResponse(ctx)
		return
	}

	ctx.File(cleanPath)
}
