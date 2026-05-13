package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zgierz/harc_quiz/backend/models"
)

func (gc GlobalController) CreateHero(ctx *gin.Context) {
	var hero models.Hero

	if err := ctx.ShouldBindJSON(&hero); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := gc.db.Create(&hero).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, hero)
}

func (gc GlobalController) GetAllHeros(ctx *gin.Context) {
	var hero []models.Hero
	gc.db.Find(&hero)
	ctx.JSON(http.StatusOK, hero)
}

func (gc GlobalController) DeleteHero(ctx *gin.Context) {
	id := ctx.Param("id")
	var hero models.Hero

	if err := gc.db.Delete(&hero, id).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Quiz and related data deleted", "success": true})
}

func (gc GlobalController) UpdateHero(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing blog id"})
		return
	}

	var hero models.Hero

	if err := gc.db.First(&hero, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	var input models.Hero

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := gc.db.Model(&hero).Updates(input).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, hero)
}
