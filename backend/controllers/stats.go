package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zgierz/harc_quiz/backend/models"
)

func (gc GlobalController) GetAllStats(ctx *gin.Context) {
	var stat models.Stat

	gc.db.Find(&stat)

	ctx.JSON(http.StatusOK, stat)
}

func (gc GlobalController) NewStat(ctx *gin.Context) {
	var stat models.Stat

	if err := ctx.ShouldBindJSON(&stat); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if stat.QuizID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	if err := gc.db.Create(&stat).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create stat"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Stat created successfully",
		"stat":    stat,
	})
}

func (gc GlobalController) CountCompletedQuizzes(ctx *gin.Context) {

	var stat models.Stat
	var count int64 = 0

	gc.db.Find(&stat).Count(&count)

	ctx.JSON(200, gin.H{"count": count})
}

func (gc GlobalController) GetWeeklyStats(ctx *gin.Context) {
	var quiz models.Quiz

	var res = gc.db.Preload("Stats").Where("stats.created_at").Find(&quiz)
	if res.Error != nil {
		ctx.JSON(500, gin.H{"error": res.Error})
	}
}

func (gc GlobalController) DeleteStat(ctx *gin.Context) {
	id := ctx.Param("id")
	var stat models.Stat

	if err := gc.db.Delete(&stat, id).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Stats data deleted", "success": true})
}
