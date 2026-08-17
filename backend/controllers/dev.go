package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/zgierz/klimson/backend/api"
)


func (controller GlobalController) BuildAndSendBinary(ctx *gin.Context) {
	chunkIndex := ctx.GetHeader("X-Chunk-Index")
	totalChunks := ctx.GetHeader("X-Total-Chunks")
	fileName := ctx.DefaultQuery("filename", "server-klimson")

	targetDir := "."
	tempFilePath := filepath.Join(targetDir, fileName+".tmp")
	finalFilePath := filepath.Join(targetDir, fileName)

	fileData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Nie udało się odczytać chunku"})
		return
	}

	f, err := os.OpenFile(tempFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		api.InternalServerErrorResponse(ctx, nil, err.Error())
		return
	}
	defer f.Close()

	if _, err := f.Write(fileData); err != nil {
		api.InternalServerErrorResponse(ctx, nil, err.Error())
		return
	}

	if chunkIndex == totalChunks {
		if err := os.Rename(tempFilePath, finalFilePath); err != nil {
			api.InternalServerErrorResponse(ctx, nil, err.Error())
			return
		}

		if err := os.Chmod(finalFilePath, 0755); err != nil {
			api.InternalServerErrorResponse(ctx, nil, err.Error())
			return
		}

		cmdBuild := exec.Command("go", "build", "-o", "server-klimson", ".")
		cmdBuild.Dir = targetDir
		if _, err := cmdBuild.CombinedOutput(); err != nil {
			api.InternalServerErrorResponse(ctx, nil, err.Error())
			return
		}

		cmdPm2 := exec.Command("pm2", "restart", "server-klimson")
		if _, err := cmdPm2.CombinedOutput(); err != nil {
			cmdPm2Start := exec.Command("pm2", "start", "server-klimson", "--name", "server-klimson")
			cmdPm2Start.Dir = targetDir
			if errStart := cmdPm2Start.Run(); errStart != nil {
				api.InternalServerErrorResponse(ctx, nil, errStart.Error())
				return
			}
		}

		ctx.JSON(http.StatusOK, gin.H{"status": "Sukces! Binarka została zaktualizowana i zrestartowana."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": fmt.Sprintf("Chunk %s z %s odebrany", chunkIndex, totalChunks)})
}

func (controller GlobalController) RegisterDevEndpoints(groupPrefix string) {
	devGroupAdmin := controller.adminPath.Group(groupPrefix)
	// devGroupPublic := controller.publicPath.Group(groupPrefix)

	devGroupAdmin.POST("/")
}