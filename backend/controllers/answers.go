package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zgierz/harc_quiz/backend/logger"
	"github.com/zgierz/harc_quiz/backend/models"
)

func (controller GlobalController) CheckAnswer(ctx *gin.Context) {
	answerId := ctx.Query("id")
	logger.ServerLog(answerId)

	var answer models.Answer
	if err := controller.db.Where("id = ?", answerId).First(&answer).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Answer not found"})
		return
	}

	if *answer.IsCorrect == true {
		ctx.JSON(200, gin.H{"id": answer.ID, "is_correct": true})
	} else {
		ctx.JSON(200, gin.H{"id": answer.ID, "is_correct": false})
	}

}
