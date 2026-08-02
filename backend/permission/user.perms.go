package permission

import "github.com/zgierz/klimson/backend/helpers"

// main users
var (
	CREATE_USER = helpers.PermissionsMetadata{
		Name:          "User Creation",
		Icon:          "mdi:user-add",
		PermissionTag: "user:create",
		Color:         "orange-500",
		Description:   "",
		Category:      USER,
	}
	UPDATE_USER = helpers.PermissionsMetadata{
		Name:          "Update the user",
		Icon:          "fa7-solid:user-pen",
		PermissionTag: "user:update",
		Color:         "green-500",
		Description:   "",
		Category:      USER,
	}
	DELETE_USER = helpers.PermissionsMetadata{
		Name:          "Delete the user",
		Icon:          "mage:user-cross-fill",
		PermissionTag: "user:delete",
		Color:         "red-500",
		Description:   "",
		Category:      USER,
	}
	GET_USER = helpers.PermissionsMetadata{
		Name:          "Get one user",
		Icon:          "mdi:user",
		PermissionTag: "user:get",
		Color:         "blue-500",
		Description:   "",
		Category:      USER,
	}
	GET_USERS_RECORDS = helpers.PermissionsMetadata{
		Name:          "Get all users",
		Icon:          "user:list",
		PermissionTag: "disk:info",
		Color:         "blue-500",
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
		Color:         "blue-500",
		Description:   "",
		Category:      USER,
	}
)

// roles
var (
	CREATE_ROLE = helpers.PermissionsMetadata{
		Name:          "Create Role",
		Icon:          "fluent:shield-task-24-regular",
		PermissionTag: "role:create",
		Color:         "blue-500",
		Description:   "Allows creating new user roles",
		Category:      USER,
	}

	GET_ROLES_RECORDS = helpers.PermissionsMetadata{
		Name:          "View All Roles",
		Icon:          "fluent:shield-task-24-regular",
		PermissionTag: "roles:read",
		Color:         "blue-500",
		Description:   "Allows viewing the list of all roles",
		Category:      USER,
	}

	GET_ROLE = helpers.PermissionsMetadata{
		Name:          "View Specific Role",
		Icon:          "fluent:shield-task-24-regular",
		PermissionTag: "role:read",
		Color:         "blue-500",
		Description:   "Allows viewing details of a specific role",
		Category:      USER,
	}

	UPDATE_ROLE = helpers.PermissionsMetadata{
		Name:          "Update Role",
		Icon:          "fluent:shield-task-24-regular",
		PermissionTag: "role:update",
		Color:         "blue-500",
		Description:   "Allows updating role details and permissions",
		Category:      USER,
	}

	DELETE_ROLE = helpers.PermissionsMetadata{
		Name:          "Delete Role",
		Icon:          "fluent:shield-task-24-regular",
		PermissionTag: "role:delete",
		Color:         "red-500",
		Description:   "Allows deleting user roles",
		Category:      USER,
	}
)
