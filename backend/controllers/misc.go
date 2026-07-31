package controllers

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
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
	khttp.SuccessResponse(ctx, gin.H{"perms": perms})
}

func (controller GlobalController) GetPermissionCategoriesList(ctx *gin.Context) {
	categories := helpers.GetAllDefinedPermissionCategories()
	khttp.SuccessResponse(ctx, gin.H{"categories": categories})
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

func getSystemInfo() (string, string) {
	arch := runtime.GOARCH

	hostInfo, err := host.Info()
	osName := "Unknown"
	if err == nil {
		osName = hostInfo.Platform + " " + hostInfo.PlatformVersion
	}

	return arch, osName
}

func (controller GlobalController) GetStorageLeftPercentage(ctx *gin.Context) {
	used, total, err := GetDiskUsage()
	arch, osName := getSystemInfo()

	if err != nil {
		khttp.InternalServerErrorResponse(ctx, nil, err.Error())
		return
	}

	if total == 0 {
		khttp.SuccessResponse(ctx, gin.H{"percentage": 0, "used": 0, "total": 0, "arch": arch, "os": osName})
		return
	}

	percentage := (float64(used) / float64(total)) * 100

	khttp.SuccessResponse(ctx, gin.H{
		"percentage": percentage,
		"used":       used,
		"total":      total,
		"label":      fmt.Sprintf("%.2f%%", percentage),
		"arch":       arch,
		"os":         osName,
	})
}

func (controller GlobalController) KlimsonFetch(ctx *gin.Context) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	stats := helpers.SystemStats{
		SystemOS:    helpers.GetUbuntuVersion(),
		GoVersion:   runtime.Version(),
		Arch:        runtime.GOARCH,
		NumCPU:      runtime.NumCPU(),
		Goroutines:  runtime.NumGoroutine(),
		ThreadCount: getThreadCountSafe(),

		MemoryAlloc: helpers.FormatBytes(m.Alloc),
		MemoryTotal: helpers.FormatBytes(m.TotalAlloc),
		MemorySys:   helpers.FormatBytes(m.Sys),
		HeapObjects: m.HeapObjects,
		NumGC:       m.NumGC,

		Uptime:     helpers.GetSystemUptime(),
		ServerTime: time.Now(),
		Timestamp:  time.Now(),
	}

	ctx.JSON(http.StatusOK, stats)
}

// Opcjonalna bezpieczna funkcja zastępcza, jeśli nie masz gotowej
func getThreadCountSafe() int {
	// W Go liczbę aktywnych wątków można też oszacować lub pobrać z debug,
	// jeśli nie jest Ci niezbędna, możesz usunąć to pole ze struktury.
	return 0
}

func (controller GlobalController) RegisterMiscEndpoints(groupPrefix string) {

	miscGroupAdmin := controller.adminPath.Group(groupPrefix)
	// miscGroupPublic := controller.publicPath.Group(groupPrefix)

	miscGroupAdmin.GET("/klimson-fetch", controller.KlimsonFetch)

}
