package helpers

import (
	"encoding/json"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

type PermissionsMetadata struct {
	Icon          string `json:"icon"`
	PermissionTag string `json:"tag"`
	Name          string `json:"name"`
	Color         string `json:"color"`
	Description   string `json:"description"`
}

var (
	permissionsMu       sync.Mutex
	permissionsRegistry = make(map[string]PermissionsMetadata)
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

func RequirePermission(meta PermissionsMetadata) gin.HandlerFunc {
	permissionsMu.Lock()
	permissionsRegistry[meta.PermissionTag] = meta
	permissionsMu.Unlock()

	return func(c *gin.Context) {

		metadataJson, _ := json.Marshal(meta.PermissionTag)

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
