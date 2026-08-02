package permission

import "github.com/zgierz/klimson/backend/helpers"

// SERVER/DEV PERMISSIONS
var (
	ADMIN = helpers.PermissionsMetadata{
		Name:          "Admin Role",
		Icon:          "lucide:binary",
		PermissionTag: "admin",
		Color:         "purple-500",
		Description:   "Has acccess to all things on the dashboard",
		Category:      SUPER,
	}
)
