package controllers

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/zgierz/klimson/backend/api"
)

// POST /dev/send?filename=server-klimson&path=/root/server/cmd
func (controller GlobalController) SendBinary(ctx *gin.Context) {
	chunkIndex := ctx.GetHeader("X-Chunk-Index")
	totalChunks := ctx.GetHeader("X-Total-Chunks")
	fileName := ctx.DefaultQuery("filename", "server-klimson")
	
	targetDir := ctx.DefaultQuery("path", ".")

	if err := os.MkdirAll(targetDir, 0755); err != nil {
		api.InternalServerErrorResponse(ctx, nil, "Directory creation error: "+err.Error())
		return
	}

	tempFilePath := filepath.Join(targetDir, fileName+".tmp")
	finalFilePath := filepath.Join(targetDir, fileName)

	fileData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		api.BadRequestResponse(ctx, nil, "Failed to read chunk data")
		return
	}

	f, err := os.OpenFile(tempFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		api.InternalServerErrorResponse(ctx, nil, "File write error: "+err.Error())
		return
	}
	defer f.Close()

	if _, err := f.Write(fileData); err != nil {
		api.InternalServerErrorResponse(ctx, nil, "Chunk write error: "+err.Error())
		return
	}

	if chunkIndex == totalChunks {
		if err := os.Rename(tempFilePath, finalFilePath); err != nil {
			api.InternalServerErrorResponse(ctx, nil, "File replacement error: "+err.Error())
			return
		}

		if err := os.Chmod(finalFilePath, 0755); err != nil {
			api.InternalServerErrorResponse(ctx, nil, "Permission grant error: "+err.Error())
			return
		}

		api.SuccessResponse(ctx, nil, "Success! The binary has been updated and restarted.")
		return
	}

	api.SuccessResponse(ctx, nil, fmt.Sprintf("Chunk %s of %s received", chunkIndex, totalChunks))
}

func (controller GlobalController) RegisterDevEndpoints(groupPrefix string) {
	devGroupAdmin := controller.adminPath.Group(groupPrefix)
	// devGroupPublic := controller.publicPath.Group(groupPrefix)

	devGroupAdmin.POST("/")
}