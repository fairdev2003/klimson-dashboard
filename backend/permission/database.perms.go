package permission

import "github.com/zgierz/klimson/backend/helpers"

var (
	GET_DB_TABLES = helpers.PermissionsMetadata{
		Name:          "Get All Available PSQL Tables",
		Icon:          "devicon-plain:postgresql",
		PermissionTag: "db:tables",
		Color:         "bg-blue-500",
		Description:   "",
		Category:      DB,
	}

	GET_DB_TABLE = helpers.PermissionsMetadata{
		Name:          "View Specific Table",
		Icon:          "devicon-plain:postgresql",
		PermissionTag: "db:table",
		Color:         "bg-blue-500",
		Description:   "",
		Category:      DB,
	}
)
