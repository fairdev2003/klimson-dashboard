package permission

import "github.com/zgierz/klimson/backend/helpers"

// SERVER/DEV PERMISSIONS
var (
	CV_RETRIEVE = helpers.PermissionsMetadata{
		Name:          "Access to cv generator path",
		Icon:          "mdi:pdf",
		PermissionTag: "server:neofetch",
		Color:         "purple-500",
		Description:   "Access generated cv",
		Category:      CV,
	}
)
