package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/zgierz/klimson/backend/api"
	"github.com/zgierz/klimson/backend/models"
	"gorm.io/gorm"
)

func (controller GlobalController) CreateContextStorage(ctx *gin.Context) {
	var context_storage models.ContextStorage

	if err := ctx.ShouldBindJSON(&context_storage); err != nil {
		api.BadRequestResponse(ctx, nil, err.Error())
		return
	}

	if err := controller.db.Create(&context_storage).Error; err != nil {
		api.InternalServerErrorResponse(ctx, nil, err.Error())
		return
	}

	api.StatusCreatedResponse(ctx, context_storage)
}

func (gc GlobalController) DeleteContextStorageRecord(ctx *gin.Context) {
	id := ctx.Param("id")
	var context_storage models.ContextStorage

	if err := gc.db.Delete(&context_storage, id).Error; err != nil {
		api.InternalServerErrorResponse(ctx, nil, err.Error())
		return
	}

	api.SuccessResponse(ctx, nil, "Stats data deleted")
}

func (gc GlobalController) GetPrivateContextStorages(ctx *gin.Context) {
	var storageItems []models.ContextStorage

	err := gc.db.Find(&storageItems).Error
	if err != nil {
		api.InternalServerErrorResponse(ctx, nil, "Failed to retrieve context storage items: "+err.Error())
		return
	}

	api.SuccessResponse(ctx, storageItems)
}

func (controller GlobalController) GetPrivateContext(ctx *gin.Context) {
	var storageItem models.ContextStorage

	key := ctx.Param("key")

	result := controller.db.First(&storageItem, "key = ?", key)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			api.NotFoundResponse(ctx)
			return
		}

		api.InternalServerErrorResponse(ctx, nil, "Failed to retrieve context storage item: "+result.Error.Error())
		return
	}

	api.SuccessResponse(ctx, storageItem)
}

func (controller GlobalController) GetPublicContextStorage(ctx *gin.Context) {
	id := ctx.Param("key")
	if id == "" || id == ":key" {
		api.BadRequestResponse(ctx, nil, "Key is required")
		return
	}

	var storage models.ContextStorage

	err := controller.db.First(&storage, "key = ?", id).Error
	if err != nil {
		api.BadRequestResponse(ctx, nil, "Context storage not found")
		return
	}

	if !*storage.Public {
		api.UnauthorizedResponse(ctx, nil, "Is not public")
		return
	}

	api.SuccessResponse(ctx, storage)
}

func (controller GlobalController) UpdateContextStorage(ctx *gin.Context) {
	key := ctx.Param("key")
	var existingItem models.ContextStorage

	result := controller.db.First(&existingItem, "key = ?", key)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			api.NotFoundResponse(ctx)
			return
		}

		api.InternalServerErrorResponse(ctx, nil, "Database error while locating item: "+result.Error.Error())
		return
	}

	var updateData models.ContextStorage
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		api.BadRequestResponse(ctx, nil, "Invalid JSON data: "+err.Error())
		return
	}

	if err := controller.db.Model(&existingItem).Updates(updateData).Error; err != nil {
		api.InternalServerErrorResponse(ctx, nil, "Failed to update context storage item: "+err.Error())
		return
	}

	api.SuccessResponse(ctx, existingItem)
}
