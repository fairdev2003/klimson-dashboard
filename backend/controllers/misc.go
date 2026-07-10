package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/disk"
	"github.com/zgierz/klimson/backend/helpers"
	"github.com/zgierz/klimson/backend/khttp"
	"github.com/zgierz/klimson/backend/models"
)

func GetDiskUsage() (uint64, uint64, error) {
	usage, err := disk.Usage("/")
	if err != nil {
		return 0, 0, err
	}
	return usage.Used, usage.Total, nil
}

func (controller GlobalController) GetPermissionsList(ctx *gin.Context) {
	perms := helpers.GetAllDefinedPermissions()

	ctx.JSON(200, perms)
}

func ListFiles(folderPath string) ([]models.ListRecord, error) {
	fullPath := "./static/uploads" + folderPath

	entries, err := os.ReadDir(fullPath)
	if err != nil {
		return nil, err
	}

	var fileList []models.ListRecord
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}

		fileList = append(fileList, models.ListRecord{
			Name:         entry.Name(),
			IsDir:        entry.IsDir(),
			ModifiedTime: info.ModTime(),
			Size:         info.Size(),
		})
	}

	return fileList, nil
}

func (controller GlobalController) TestRedirect(ctx *gin.Context) {
	ctx.Redirect(http.StatusTemporaryRedirect, "/login")
}

func (controller GlobalController) LegalReasonsTest(ctx *gin.Context) {
	ctx.JSON(http.StatusUnavailableForLegalReasons, gin.H{"message": "451"})
}

func (controller GlobalController) GetStorageLeftPercentage(ctx *gin.Context) {
	used, total, err := GetDiskUsage()
	if err != nil {
		khttp.InternalServerErrorResponse(ctx, nil, err.Error())
		return
	}

	if total == 0 {
		khttp.SuccessResponse(ctx, gin.H{"percentage": 0, "used": 0, "total": 0})
		return
	}

	percentage := (float64(used) / float64(total)) * 100

	khttp.SuccessResponse(ctx, gin.H{
		"percentage": percentage,
		"used":       used,
		"total":      total,
		"label":      fmt.Sprintf("%.2f%%", percentage),
	})
}
