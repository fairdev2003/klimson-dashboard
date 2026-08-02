package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zgierz/klimson/backend/helpers"
	"github.com/zgierz/klimson/backend/khttp"
	"github.com/zgierz/klimson/backend/models"
	"github.com/zgierz/klimson/backend/permission"
)

// 1. CREATE (NewUser już masz, tutaj tylko wersja z Preloadem dla spójności)
func (gc GlobalController) NewUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		khttp.BadRequestResponse(ctx, nil, "Invalid json")
		return
	}

	if err := gc.db.Create(&user).Error; err != nil {
		khttp.BadRequestResponse(ctx, nil, err.Error())
		return
	}

	gc.db.Preload("Role").First(&user, user.ID)
	khttp.SuccessResponse(ctx, user, "User created successfully")
}

func (gc GlobalController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.User

	if err := gc.db.Preload("Role").First(&user, id).Error; err != nil {
		khttp.BadRequestResponse(ctx, nil, "User not found")
		return
	}
	khttp.SuccessResponse(ctx, gin.H{"user": user}, "User details")
}

func (gc GlobalController) ListUsers(ctx *gin.Context) {
	var users []models.User
	gc.db.Preload("Role").Find(&users)
	khttp.SuccessResponse(ctx, gin.H{"users": users}, "Users list")
}

func (gc GlobalController) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.User

	if err := gc.db.First(&user, id).Error; err != nil {
		khttp.BadRequestResponse(ctx, nil, "User not found")
		return
	}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		khttp.BadRequestResponse(ctx, nil, "Invalid json")
		return
	}

	gc.db.Save(&user)
	khttp.SuccessResponse(ctx, nil, "User updated successfully")
}

func (gc GlobalController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := gc.db.Delete(&models.User{}, id).Error; err != nil {
		khttp.BadRequestResponse(ctx, nil, "Failed to delete user")
		return
	}
	khttp.SuccessResponse(ctx, nil, "User deleted successfully")
}

func (gc GlobalController) GetYourself(ctx *gin.Context) {
	isRoot := ctx.GetBool("isRoot")
	if isRoot {
		claims, exists := ctx.Get("claims")
		if !exists {
			return
		}

		khttp.SuccessResponse(ctx, gin.H{"claims": claims})
		return
	}

}

func (gc GlobalController) RegisterUserController() {
	users := gc.adminPath.Group("/users")
	{
		users.POST("/new-user", helpers.RequirePermission(permission.CREATE_USER), gc.NewUser)
		users.GET("/get-users", helpers.RequirePermission(permission.GET_USERS_RECORDS), gc.ListUsers)
		users.GET("/get-user/:id", helpers.RequirePermission(permission.GET_USER), gc.GetUser)
		users.PUT("/update-user/:id", helpers.RequirePermission(permission.UPDATE_USER), gc.UpdateUser)
		users.DELETE("/delete-user/:id", helpers.RequirePermission(permission.DELETE_USER), gc.DeleteUser)
		users.GET("/me", gc.GetYourself)

	}
}
