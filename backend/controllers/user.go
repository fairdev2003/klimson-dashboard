package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zgierz/klimson/backend/api"
	"github.com/zgierz/klimson/backend/helpers"
	"github.com/zgierz/klimson/backend/models"
	"github.com/zgierz/klimson/backend/permission"
)

// 1. CREATE (NewUser już masz, tutaj tylko wersja z Preloadem dla spójności)
func (gc GlobalController) NewUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		api.BadRequestResponse(ctx, nil, "Invalid json")
		return
	}

	if err := gc.db.Create(&user).Error; err != nil {
		api.BadRequestResponse(ctx, nil, err.Error())
		return
	}

	gc.db.Preload("Role").First(&user, user.ID)
	api.SuccessResponse(ctx, user, "User created successfully")
}

func (gc GlobalController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.User

	if err := gc.db.Preload("Role").First(&user, id).Error; err != nil {
		api.BadRequestResponse(ctx, nil, "User not found")
		return
	}
	api.SuccessResponse(ctx, gin.H{"user": user}, "User details")
}

func (gc GlobalController) ListUsers(ctx *gin.Context) {
	var users []models.User
	gc.db.Preload("Role").Find(&users)
	api.SuccessResponse(ctx, gin.H{"users": users}, "Users list")
}

func (gc GlobalController) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.User

	if err := gc.db.First(&user, id).Error; err != nil {
		api.BadRequestResponse(ctx, nil, "User not found")
		return
	}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		api.BadRequestResponse(ctx, nil, "Invalid json")
		return
	}

	gc.db.Save(&user)
	api.SuccessResponse(ctx, nil, "User updated successfully")
}

func (gc GlobalController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := gc.db.Delete(&models.User{}, id).Error; err != nil {
		api.BadRequestResponse(ctx, nil, "Failed to delete user")
		return
	}
	api.SuccessResponse(ctx, nil, "User deleted successfully")
}

func (gc GlobalController) GetYourself(ctx *gin.Context) {
	isRoot := ctx.GetBool("isRoot")
	if isRoot {
		claims, exists := ctx.Get("claims")
		if !exists {
			return
		}

		api.SuccessResponse(ctx, gin.H{"claims": claims})
		return
	}

}

type PermissionInput struct {
	Name string `json:"name"`
}

type RoleInput struct {
	Name        string              `json:"name" binding:"required"`
	Color       string              `json:"color"`
	Icon        string              `json:"icon"`
	Permissions []models.Permission `json:"permissions"`
}

func (gc GlobalController) CreateRole(ctx *gin.Context) {
	var input RoleInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		api.BadRequestResponse(ctx, nil, "JSON Error: "+err.Error())
		return
	}

	var permissions []models.Permission

	for _, pInput := range input.Permissions {
		var perm models.Permission
		gc.db.Where(models.Permission{Name: pInput.Name}).FirstOrCreate(&perm)
		permissions = append(permissions, perm)
	}

	role := models.Role{
		Name:        input.Name,
		Color:       input.Color,
		Icon:        input.Icon,
		Permissions: permissions,
	}

	if result := gc.db.Create(&role); result.Error != nil {
		api.InternalServerErrorResponse(ctx, nil, "Error during creating role: "+result.Error.Error())
		return
	}

	gc.db.Preload("Permissions").First(&role, role.ID)

	api.SuccessResponse(ctx, gin.H{"data": role}, "Role created successfully.")
}

func (gc GlobalController) GetRoles(ctx *gin.Context) {
	var roles []models.Role
	if result := gc.db.Preload("Permissions").Find(&roles); result.Error != nil {
		api.InternalServerErrorResponse(ctx, nil, "Error during fetching the role: "+result.Error.Error())
		return
	}

	api.SuccessResponse(ctx, gin.H{"data": roles})
}

func (gc GlobalController) GetRole(ctx *gin.Context) {
	id := ctx.Param("id")
	var role models.Role

	if result := gc.db.Preload("Permissions").First(&role, id); result.Error != nil {
		api.NotFoundResponse(ctx)
		return
	}

	api.SuccessResponse(ctx, gin.H{"data": role}, "Successfully fetched the role.")

}

func (gc GlobalController) UpdateRole(ctx *gin.Context) {
	id := ctx.Param("id")
	var role models.Role

	if result := gc.db.First(&role, id); result.Error != nil {
		api.NotFoundResponse(ctx)
		return
	}

	var input RoleInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		api.BadRequestResponse(ctx, nil, "JSON Erorr: "+err.Error())
		return
	}

	var permissions []models.Permission
	if len(input.Permissions) > 0 {
		gc.db.Find(&permissions, input.Permissions)
	}

	role.Name = input.Name
	role.Color = input.Color
	gc.db.Save(&role)

	gc.db.Model(&role).Association("Permissions").Replace(&permissions)

	api.SuccessResponse(ctx, nil, "Roles was updated successfully.")
}

func (gc GlobalController) DeleteRole(ctx *gin.Context) {
	id := ctx.Param("id")
	var role models.Role

	if result := gc.db.First(&role, id); result.Error != nil {
		api.NotFoundResponse(ctx)
		return
	}

	if result := gc.db.Where("role_id = ?", role.ID).Delete(&models.Permission{}); result.Error != nil {
		api.InternalServerErrorResponse(ctx, nil, "Error during deleting permissions: "+result.Error.Error())
		return
	}

	if result := gc.db.Delete(&role); result.Error != nil {
		api.InternalServerErrorResponse(ctx, nil, "Error during deleting the role: "+result.Error.Error())
		return
	}

	api.SuccessResponse(ctx, nil, "Role and its permissions are successfully deleted.")
}

func (gc GlobalController) RegisterUserController() {
	users := gc.adminPath.Group("/users")
	users.POST("/new-user", helpers.RequirePermission(permission.CREATE_USER), gc.NewUser)
	users.GET("/get-users", helpers.RequirePermission(permission.GET_USERS_RECORDS), gc.ListUsers)
	users.GET("/get-user/:id", helpers.RequirePermission(permission.GET_USER), gc.GetUser)
	users.PUT("/update-user/:id", helpers.RequirePermission(permission.UPDATE_USER), gc.UpdateUser)
	users.DELETE("/delete-user/:id", helpers.RequirePermission(permission.DELETE_USER), gc.DeleteUser)
	users.GET("/me", gc.GetYourself)

}

func (gc GlobalController) RegisterRoleController() {
	roles := gc.adminPath.Group("/users/roles")
	roles.POST("/new-role", helpers.RequirePermission(permission.CREATE_ROLE), gc.CreateRole)
	roles.GET("/get-roles", helpers.RequirePermission(permission.GET_ROLES_RECORDS), gc.GetRoles)
	roles.GET("/get-role/:id", helpers.RequirePermission(permission.GET_ROLE), gc.GetRole)
	roles.PUT("/update-role/:id", helpers.RequirePermission(permission.UPDATE_ROLE), gc.UpdateRole)
	roles.DELETE("/delete-role/:id", helpers.RequirePermission(permission.DELETE_ROLE), gc.DeleteRole)
}
