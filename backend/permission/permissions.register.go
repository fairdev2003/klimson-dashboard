package permission

import "github.com/zgierz/klimson/backend/helpers"

// categories
var (
	STORAGE_V2     = "Storage V2"
	PUBLIC_STORAGE = "Public Storage"
	MISC           = "miscellaneous"
)

// regsiter
var (
	V2_CREATE_FOLDER = helpers.PermissionsMetadata{
		Name:          "Create Private Folder (v2)",
		Icon:          "material-symbols:folder",
		PermissionTag: "storage:private:create-folder",
		Color:         "bg-blue-500",
		Description:   "",
		Category:      STORAGE_V2,
	}
	STORAGE_CREATE_FOLDER = helpers.PermissionsMetadata{
		Name:          "Create Private Folder",
		Icon:          "material-symbols:folder",
		PermissionTag: "storage:create-folder",
		Color:         "bg-blue-500",
		Description:   "",
		Category:      PUBLIC_STORAGE,
	}

	STORAGE_LATEST = helpers.PermissionsMetadata{
		Name:          "View Latest File",
		Icon:          "mdi:latest",
		PermissionTag: "storage:latest",
		Color:         "bg-blue-500",
		Description:   "",
		Category:      PUBLIC_STORAGE,
	}

	DISK_INFO = helpers.PermissionsMetadata{
		Name:          "Disk Info",
		Icon:          "vaadin:harddrive",
		PermissionTag: "disk:info",
		Color:         "bg-blue-500",
		Description:   "",
		Category:      MISC,
	}
)
