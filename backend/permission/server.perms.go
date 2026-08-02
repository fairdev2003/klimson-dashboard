package permission

import "github.com/zgierz/klimson/backend/helpers"

// SERVER/DEV PERMISSIONS
var (
	SERVER_BINARY = helpers.PermissionsMetadata{
		Name:          "Send Binary Startup File",
		Icon:          "lucide:binary",
		PermissionTag: "dev:binary-send",
		Color:         "bg-blue-500",
		Description:   "",
		Category:      SERVER,
	}
)
