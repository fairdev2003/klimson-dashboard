package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zgierz/harc_quiz/backend/khttp"
	"github.com/zgierz/harc_quiz/backend/models"
	"gorm.io/gorm"
)

func (controller GlobalController) CreateContextStorage(ctx *gin.Context) {
	var context_storage models.ContextStorage

	if err := ctx.ShouldBindJSON(&context_storage); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.db.Create(&context_storage).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, context_storage)
}

func (gc GlobalController) DeleteContextStorageRecord(ctx *gin.Context) {
	id := ctx.Param("id")
	var context_storage models.ContextStorage

	if err := gc.db.Delete(&context_storage, id).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Stats data deleted", "success": true})
}

func (gc GlobalController) GetPrivateContextStorages(ctx *gin.Context) {
	var storageItems []models.ContextStorage

	err := gc.db.Find(&storageItems).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to retrieve context storage items",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, storageItems)
}

func (controller GlobalController) GetPrivateContext(ctx *gin.Context) {
	var storageItem models.ContextStorage

	key := ctx.Param("key")

	result := controller.db.First(&storageItem, "key = ?", key)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "error",
				"message": "Context storage item not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to retrieve context storage item",
			"details": result.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, storageItem)
}

func (controller GlobalController) GetPublicContextStorage(ctx *gin.Context) {
	id := ctx.Param("key")
	if id == "" || id == ":key" {
		khttp.BadRequestResponse(ctx, nil, "Key is required")
		return
	}

	var storage models.ContextStorage

	err := controller.db.First(&storage, "key = ?", id).Error
	if err != nil {
		khttp.BadRequestResponse(ctx, nil, "Context storage not found")
		return
	}

	if !*storage.Public {
		khttp.UnauthorizedResponse(ctx, nil, "Is not public")
		return
	}

	ctx.JSON(200, storage)
}

func (controller GlobalController) UpdateContextStorage(ctx *gin.Context) {
	key := ctx.Param("key")
	var existingItem models.ContextStorage

	result := controller.db.First(&existingItem, "key = ?", key)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "error",
				"message": "Context storage item not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Database error while locating item",
			"details": result.Error.Error(),
		})
		return
	}

	var updateData models.ContextStorage
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid JSON data",
			"details": err.Error(),
		})
		return
	}

	if err := controller.db.Model(&existingItem).Updates(updateData).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to update context storage item",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, existingItem)
}
