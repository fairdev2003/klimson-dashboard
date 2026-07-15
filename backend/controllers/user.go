package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zgierz/klimson/backend/khttp"
	"github.com/zgierz/klimson/backend/models"
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
	khttp.SuccessResponse(ctx, user, "User details")
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
	khttp.SuccessResponse(ctx, user, "User updated successfully")
}

func (gc GlobalController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := gc.db.Delete(&models.User{}, id).Error; err != nil {
		khttp.BadRequestResponse(ctx, nil, "Failed to delete user")
		return
	}
	khttp.SuccessResponse(ctx, nil, "User deleted successfully")
}

func (gc GlobalController) RegisterUserController() {
	users := gc.adminPath.Group("/users")
	{
		users.POST("/new-user", gc.NewUser)
		users.GET("/get-users", gc.ListUsers)
		users.GET("/get-user/:id", gc.GetUser)
		users.PUT("/update-user/:id", gc.UpdateUser)
		users.DELETE("/delete-user/:id", gc.DeleteUser)
	}
}
