package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zgierz/harc_quiz/backend/models"
	"gorm.io/gorm"
)

func (controller GlobalController) GetPublicQuizes(ctx *gin.Context) {
	var quizzes []models.Quiz

	if err := controller.db.
		Preload("Questions").
		Preload("Stats").
		Find(&quizzes).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch quizzes"})
		return
	}

	ctx.JSON(http.StatusOK, quizzes)
}

type ImageForm struct {
	ImageUrl string `json:"image_url"`
}

func (controller GlobalController) UpdateQuizImage(ctx *gin.Context) {
	quizID := ctx.Param("id")
	var imageFormData ImageForm
	if err := ctx.ShouldBindJSON(&imageFormData); err != nil {
		ctx.JSON(400, gin.H{"error": "Nieprawidłowe dane: " + err.Error()})
		return
	}

	result := controller.db.Model(&models.Quiz{}).Where("id = ?", quizID).Updates(imageFormData)

	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": "Błąd podczas aktualizacji bazy"})
		return
	}

	if result.RowsAffected == 0 {
		ctx.JSON(404, gin.H{"error": "Nie znaleziono quizu o podanym ID"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Zdjecie zostało zaaktualizowane", "id": quizID})
}

type SettingsUpdateModel struct {
	Public *bool `json:"public"`
}

func (controller GlobalController) UpdateQuizSettings(ctx *gin.Context) {
	quizID := ctx.Param("id")
	var settingsUpdateData SettingsUpdateModel
	if err := ctx.ShouldBindJSON(&settingsUpdateData); err != nil {
		ctx.JSON(400, gin.H{"error": "Nieprawidłowe dane: " + err.Error()})
		return
	}

	result := controller.db.Model(&models.Quiz{}).Where("id = ?", quizID).Updates(settingsUpdateData)

	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": "Błąd podczas aktualizacji bazy"})
		return
	}

	if result.RowsAffected == 0 {
		ctx.JSON(404, gin.H{"error": "Nie znaleziono quizu o podanym ID"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Informacje zostały zaktualizowane", "id": quizID})

}

type BasicInfoUpdateModel struct {
	Title       string `json:"title" binding:"required,min=3"`
	Description string `json:"description"`
	Difficulty  string `json:"difficulty"`
	Author      string `json:"author"`
}

func (controller GlobalController) UpdateBasicInfo(ctx *gin.Context) {
	quizID := ctx.Param("id")
	var updateData BasicInfoUpdateModel

	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(400, gin.H{"error": "Nieprawidłowe dane: " + err.Error()})
		return
	}

	result := controller.db.Model(&models.Quiz{}).Where("id = ?", quizID).Updates(updateData)

	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": "Błąd podczas aktualizacji bazy"})
		return
	}

	if result.RowsAffected == 0 {
		ctx.JSON(404, gin.H{"error": "Nie znaleziono quizu o podanym ID"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Informacje zostały zaktualizowane", "id": quizID})
}

func (controller GlobalController) SaveBasicInfo(ctx *gin.Context) {
	var input BasicInfoUpdateModel
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": "Błędne dane"})
		return
	}

	quiz := models.Quiz{
		Title:       input.Title,
		Description: input.Description,
		Author:      input.Author,
		Difficulty:  input.Difficulty,
	}

	result := controller.db.Save(&quiz)

	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": "Błąd bazy"})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Zapisano pomyślnie",
		"id":      quiz.ID,
	})
}

func (controller GlobalController) GetAdminQuestions(ctx *gin.Context) {
	var quiz []models.Question

	err := controller.db.Preload("Questions").Find(&quiz).Order("id DESC")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	ctx.JSON(200, quiz)
}

func (controller GlobalController) UpdateOneField(ctx *gin.Context) {

	// quiz id
	id := ctx.Param("id")

	// We use a map to capture the dynamic value from JSON
	var input map[string]any
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	// Example input: {"key": "done_count", "value": 42}
	key, keyExists := input["key"].(string)
	val, valExists := input["value"]

	if !keyExists || !valExists {
		ctx.JSON(400, gin.H{"error": "Missing field or value"})
		return
	}

	// GORM handles the type conversion to SQL automatically
	result := controller.db.Model(&models.Question{}).
		Where("id = ?", id).
		Update(key, val)

	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Updated successfully"})
}

func (controller GlobalController) GetAdminQuizzes(ctx *gin.Context) {
	var quiz []models.Quiz

	err := controller.db.Preload("Questions.Answers").Preload("Stats").Find(&quiz)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	ctx.JSON(200, quiz)
}

func (controller GlobalController) GetAdminQuiz(ctx *gin.Context) {
	var quiz models.Quiz
	idStr := ctx.Param("id")
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing id"})
		return
	}

	var id uint64
	if parsed, err := strconv.ParseUint(idStr, 10, 64); err == nil {
		id = parsed
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err := controller.db.
		Preload("Questions").
		Preload("Questions.Answers").
		First(&quiz, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "quiz not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, quiz)
}

func (controller GlobalController) GetPublicQuiz(ctx *gin.Context) {
	var quiz models.Quiz
	idStr := ctx.Query("id")
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing id"})
		return
	}

	// opcjonalnie konwertujemy id do int (bezpośrednie First(&quiz, id) też działa)
	var id uint64
	if parsed, err := strconv.ParseUint(idStr, 10, 64); err == nil {
		id = parsed
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	// Debug -> włącz log SQL tymczasowo, żeby zobaczyć co GORM robi
	db := controller.db.Debug()

	err := db.
		Where("public = ?", true). // <-- tylko publiczne quizy
		Preload("Questions", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "quiz_id", "content", "image_url", "type", "time_limit")
		}).
		Preload("Questions.Answers", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "question_id", "content", "is_correct") // whitelist pól
		}).
		First(&quiz, id).Error // pobierz rekord po PK

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "quiz not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, quiz)
}

func (controller GlobalController) CreateQuiz(ctx *gin.Context) {
	var quiz models.Quiz

	if err := ctx.ShouldBindJSON(&quiz); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.db.Create(&quiz).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, quiz)
}

func (controller GlobalController) UpdateQuiz(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing quiz id"})
		return
	}

	var quiz models.Quiz

	// Sprawdź czy quiz istnieje
	if err := controller.db.First(&quiz, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Quiz not found"})
		return
	}

	// Payload – tylko pola quizu
	var input models.Quiz
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update tylko pól quizu
	if err := controller.db.
		Model(&quiz).
		Updates(&input).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update quiz"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Quiz updated successfully",
	})
}

func (controller GlobalController) DeleteQuiz(ctx *gin.Context) {
	id := ctx.Param("id")

	var quiz models.Quiz
	if err := controller.db.First(&quiz).Where("id = ?", id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Quiz not found"})
		return
	}

	// Usuń odpowiedzi
	for _, q := range quiz.Questions {
		controller.db.Delete(&models.Answer{}, "question_id = ?", q.ID)
	}

	// Usuń pytania
	controller.db.Delete(&models.Question{}, "quiz_id = ?", id)

	// Usuń quiz
	if err := controller.db.Delete(&models.Quiz{}, id).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Quiz and related data deleted", "success": true})
}
