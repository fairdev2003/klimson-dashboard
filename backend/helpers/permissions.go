package helpers

import (
	"encoding/json"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/zgierz/klimson/backend/logger"
)

type PermissionsMetadata struct {
	Icon          string `json:"icon"`
	PermissionTag string `json:"tag"`
	Name          string `json:"name"`
	Color         string `json:"color"`
	Description   string `json:"description"`
	Category      string `json:"category"`
}

var (
	permissionsMu       sync.Mutex
	permissionsRegistry = make(map[string]PermissionsMetadata)
	categoriesList      = []string{}
	uniqueCategories    = make(map[string]struct{})
)

func GetAllDefinedPermissions() []PermissionsMetadata {
	permissionsMu.Lock()
	defer permissionsMu.Unlock()

	list := make([]PermissionsMetadata, 0, len(permissionsRegistry))
	for _, metadata := range permissionsRegistry {
		list = append(list, metadata)
	}
	return list
}

func GetAllDefinedPermissionCategories() []string {
	permissionsMu.Lock()
	defer permissionsMu.Unlock()

	list := make([]string, 0, len(categoriesList))
	for _, categoryName := range categoriesList {
		list = append(list, categoryName)
	}
	return list
}

func appendCategories(meta PermissionsMetadata) {
	uniqueCategories[meta.Category] = struct{}{}

	var categoriesList []string
	for category := range uniqueCategories {
		categoriesList = append(categoriesList, category)
	}
}

func appendPermissions(meta PermissionsMetadata) {
	permissionsMu.Lock()
	permissionsRegistry[meta.PermissionTag] = meta
	permissionsMu.Unlock()
}

func RequirePermission(meta PermissionsMetadata) gin.HandlerFunc {
	// appending permissions for later fetch from the app state
	appendPermissions(meta)

	// creating categories list based on added permission via **RequirePermission** function
	appendCategories(meta)

	return func(c *gin.Context) {

		metadataJson, _ := json.Marshal(meta.PermissionTag)
		logger.ServerLog("Permission Tag:", meta.PermissionTag)

		c.Writer.Header().Set("X-Required-Permission", string(metadataJson))

		c.Writer.Header().Set("Access-Control-Expose-Headers", "X-Required-Permission")

		if c.GetBool("isRoot") {
			c.Next()
			return
		}

		userHasPermissionsRaw := c.GetString("permissions")
		userPerms := strings.Split(userHasPermissionsRaw, ",")

		var requestedID string
		if strings.Contains(meta.PermissionTag, "$") {
			requestedID = c.Param("id")
		}

		hasAccess := false
		for _, p := range userPerms {
			cleanPerm := strings.TrimSpace(p)

			// 1. Super-user wildcard
			if cleanPerm == "*" {
				hasAccess = true
				break
			}

			if strings.HasSuffix(cleanPerm, ":*") {
				prefix := strings.TrimSuffix(cleanPerm, "*")
				if strings.HasPrefix(meta.PermissionTag, prefix) {
					hasAccess = true
					break
				}
			}

			patternToCheck := meta.PermissionTag
			if requestedID != "" {
				patternToCheck = strings.Replace(meta.PermissionTag, "$", requestedID, 1)
			}

			if cleanPerm == patternToCheck {
				hasAccess = true
				break
			}
		}

		if hasAccess {
			c.Next()
		} else {
			c.AbortWithStatusJSON(403, gin.H{
				"error": "Brak uprawnień: " + meta.Name,
			})
		}
	}
}

const (
	IMAGE_UPLOAD = "image:upload"
)
