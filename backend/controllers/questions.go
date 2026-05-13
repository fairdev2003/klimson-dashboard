package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zgierz/harc_quiz/backend/models"
	"gorm.io/gorm"
)

func (controller GlobalController) GetPublicQuestions(ctx *gin.Context) {
	quizID, err := strconv.Atoi(ctx.Param("quiz_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quiz id"})
		return
	}

	var questions []models.PublicQuestion
	if err := controller.db.Model(&models.Question{}).
		Where("quiz_id = ?", quizID).
		Select("id, quiz_id, content, image_url, type").
		Preload("Answers", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, question_id, content")
		}).
		Find(&questions).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

}

func (controller GlobalController) GetQuestionByID(ctx *gin.Context) {
	id := ctx.Param("id")
	var question models.PublicQuestion

	if err := controller.db.Model(&models.Question{}).
		Select("id, quiz_id, content, image_url, type").
		Preload("Answers", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, question_id, content")
		}).
		First(&question, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}

	ctx.JSON(http.StatusOK, question)
}

func (controller GlobalController) GetQuestion(ctx *gin.Context) {

}

func (controller GlobalController) CreateQuestion(ctx *gin.Context) {
	quizID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quiz id"})
		return
	}

	var question models.Question
	if err := ctx.ShouldBindJSON(&question); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	question.QuizID = uint(quizID)

	if err := controller.db.Create(&question).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, question)
}

func (controller GlobalController) UpdateQuestions(ctx *gin.Context) {
	var questions []models.Question

	if err := ctx.ShouldBindJSON(&questions); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := controller.db.Session(&gorm.Session{FullSaveAssociations: true}).Select("*").Save(&questions).Error

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Błąd podczas zapisu pytań"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Pytania zaktualizowane pomyślnie"})
}

func (controller GlobalController) UpdateQuestion(ctx *gin.Context) {
	id := ctx.Param("id")

	var question models.Question
	if err := controller.db.Preload("Answers").First(&question, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}

	var input models.Question
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := controller.db.Begin()

	if err := tx.Model(&question).Updates(&input).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := tx.Where("question_id = ?", question.ID).
		Delete(&models.Answer{}).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, answer := range input.Answers {
		answer.ID = 0
		answer.QuestionID = question.ID

		answer.QuestionID = question.ID
		if err := tx.Create(&answer).Error; err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	tx.Commit()

	controller.db.Preload("Answers").First(&question, question.ID)
	ctx.JSON(http.StatusOK, question)
}

func (controller GlobalController) DeleteQuestion(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := controller.db.Delete(&models.Question{}, id).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Question deleted"})
}
