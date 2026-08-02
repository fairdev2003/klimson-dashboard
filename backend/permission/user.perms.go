package permission

import "github.com/zgierz/klimson/backend/helpers"

// main users
var (
	CREATE_USER = helpers.PermissionsMetadata{
		Name:          "User Creation",
		Icon:          "mdi:user-add",
		PermissionTag: "user:create",
		Color:         "bg-blue-500",
		Description:   "",
		Category:      USER,
	}
	UPDATE_USER = helpers.PermissionsMetadata{
		Name:          "Update the user",
		Icon:          "fa7-solid:user-pen",
		PermissionTag: "user:update",
		Color:         "bg-blue-500",
		Description:   "",
		Category:      USER,
	}
	DELETE_USER = helpers.PermissionsMetadata{
		Name:          "Delete the user",
		Icon:          "mage:user-cross-fill",
		PermissionTag: "user:delete",
		Color:         "bg-blue-500",
		Description:   "",
		Category:      USER,
	}
	GET_USER = helpers.PermissionsMetadata{
		Name:          "Get one user",
		Icon:          "mdi:user",
		PermissionTag: "user:get",
		Color:         "bg-blue-500",
		Description:   "",
		Category:      USER,
	}
	GET_USERS_RECORDS = helpers.PermissionsMetadata{
		Name:          "Get all users",
		Icon:          "user:list",
		PermissionTag: "disk:info",
		Color:         "bg-blue-500",
		Description:   "",
		Category:      USER,
	}
)

// permissions
var (
	GET_PERMISSION_REGISTER = helpers.PermissionsMetadata{
		Name:          "Get Permission register",
		Icon:          "bitcoin-icons:two-keys-filled",
		PermissionTag: "roles:permissions:register",
		Color:         "bg-blue-500",
		Description:   "",
		Category:      USER,
	}
)

// roles
var ()
