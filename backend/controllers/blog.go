package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zgierz/harc_quiz/backend/logger"
	"github.com/zgierz/harc_quiz/backend/models"
)

func (gc GlobalController) CreateBlog(ctx *gin.Context) {
	var blog models.Blog

	if err := ctx.ShouldBindJSON(&blog); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := gc.db.Create(&blog).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, blog)
}

func (gc GlobalController) GetAllBlogs(ctx *gin.Context) {
	var blog []models.Blog
	gc.db.Find(&blog)
	ctx.JSON(http.StatusOK, blog)
}

func (gc GlobalController) DeleteBlog(ctx *gin.Context) {
	id := ctx.Param("id")
	var blog models.Blog

	if err := gc.db.Delete(&blog, id).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Blog and related data deleted", "success": true})
}

func (gc GlobalController) UpdateBlog(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing blog id"})
		return
	}

	var blog models.Blog

	if err := gc.db.First(&blog, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	var input models.Blog
	logger.ServerLog(input.Public)
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := gc.db.Model(&blog).Updates(input).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, blog)
}
