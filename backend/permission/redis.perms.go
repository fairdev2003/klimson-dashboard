package permission

import "github.com/zgierz/klimson/backend/helpers"

// REDIS PERMISSIONS
var (
	SET_REDIS_KEY = helpers.PermissionsMetadata{
		Name:          "Get All Users",
		Icon:          "simple-icons:redis",
		PermissionTag: "redis:set",
		Color:         "red-500",
		Description:   "",
		Category:      REDIS,
	}
	GET_REDIS_KEY = helpers.PermissionsMetadata{
		Name:          "Get Specific Redis Private Key",
		Icon:          "simple-icons:redis",
		PermissionTag: "redis:get",
		Color:         "red-500",
		Description:   "",
		Category:      REDIS,
	}

	DEL_REDIS_KEY = helpers.PermissionsMetadata{
		Name:          "Delete Specific Redis Key",
		Icon:          "simple-icons:redis",
		PermissionTag: "redis:del",
		Color:         "red-500",
		Description:   "",
		Category:      REDIS,
	}
	REDIS_KEY_INFO = helpers.PermissionsMetadata{
		Name:          "Get Basic Information About Key",
		Icon:          "simple-icons:redis",
		PermissionTag: "redis:key-info",
		Color:         "red-500",
		Description:   "",
		Category:      REDIS,
	}
)
