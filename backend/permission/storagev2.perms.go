package permission

import "github.com/zgierz/klimson/backend/helpers"

// PRIVATE STORAGE PERMISSIONS
var (
	V2_CREATE_FOLDER = helpers.PermissionsMetadata{
		Name:          "Create Private Folder (v2)",
		Icon:          "material-symbols:folder",
		PermissionTag: "storage:private:create-folder",
		Color:         "blue-500",
		Description:   "",
		Category:      STORAGE_V2,
	}
	V2_GET_FILE = helpers.PermissionsMetadata{
		Name:          "Create Private Folder (v2)",
		Icon:          "material-symbols:folder",
		PermissionTag: "storage:private:get-file",
		Color:         "blue-500",
		Description:   "",
		Category:      STORAGE_V2,
	}
	V2_NEW_FILE = helpers.PermissionsMetadata{
		Name:          "Create Private Folder (v2)",
		Icon:          "material-symbols:folder",
		PermissionTag: "storage:private:new-file",
		Color:         "blue-500",
		Description:   "",
		Category:      STORAGE_V2,
	}
)
