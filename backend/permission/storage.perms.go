package permission

import "github.com/zgierz/klimson/backend/helpers"

// STORAGE PERMISSIONS
var (
	STORAGE_CREATE_FOLDER = helpers.PermissionsMetadata{
		Name:          "Create Private Folder",
		Icon:          "material-symbols:folder",
		PermissionTag: "storage:create-folder",
		Color:         "bg-blue-500",
		Description:   "Allows creating new folders in storage",
		Category:      PUBLIC_STORAGE,
	}

	STORAGE_UPLOAD_FILE = helpers.PermissionsMetadata{
		Name:          "Upload File",
		Icon:          "material-symbols:upload-file",
		PermissionTag: "storage:upload-file",
		Color:         "bg-blue-500",
		Description:   "Allows uploading files to storage",
		Category:      PUBLIC_STORAGE,
	}

	STORAGE_DELETE = helpers.PermissionsMetadata{
		Name:          "Delete File or Folder",
		Icon:          "material-symbols:delete",
		PermissionTag: "storage:delete",
		Color:         "bg-red-500",
		Description:   "Allows deleting files and folders",
		Category:      PUBLIC_STORAGE,
	}

	STORAGE_RENAME = helpers.PermissionsMetadata{
		Name:          "Rename Item",
		Icon:          "material-symbols:drive-file-rename-outline",
		PermissionTag: "storage:rename",
		Color:         "bg-yellow-500",
		Description:   "Allows renaming files and folders",
		Category:      PUBLIC_STORAGE,
	}

	STORAGE_EDIT_FILE = helpers.PermissionsMetadata{
		Name:          "Edit Text File",
		Icon:          "material-symbols:edit-document",
		PermissionTag: "storage:edit-file",
		Color:         "bg-green-500",
		Description:   "Allows modifying contents of text files",
		Category:      PUBLIC_STORAGE,
	}

	STORAGE_NEW_FILE = helpers.PermissionsMetadata{
		Name:          "Create New File",
		Icon:          "material-symbols:note-add",
		PermissionTag: "storage:new-file",
		Color:         "bg-blue-500",
		Description:   "Allows creating empty text files",
		Category:      PUBLIC_STORAGE,
	}

	STORAGE_LATEST = helpers.PermissionsMetadata{
		Name:          "View Latest File",
		Icon:          "mdi:latest",
		PermissionTag: "storage:latest",
		Color:         "bg-blue-500",
		Description:   "Allows viewing latest storage records",
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
