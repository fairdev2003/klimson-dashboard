package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zgierz/harc_quiz/backend/helpers"
	"github.com/zgierz/harc_quiz/backend/logger"
	"github.com/zgierz/harc_quiz/backend/models"
	"golang.org/x/crypto/bcrypt"
)

func (controller GlobalController) GetPermissionsList(ctx *gin.Context) {
	perms := helpers.GetAllDefinedPermissions()

	ctx.JSON(200, perms)
}

func (controller GlobalController) GetContributors(c *gin.Context) {
	var contributors []models.Contributor

	if err := controller.db.Find(&contributors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve contributors"})
		return
	}

	c.JSON(http.StatusOK, contributors)
}

func (controller GlobalController) SwitchContributorBlock(c *gin.Context) {
	var contributor models.Contributor
	contributorId := c.Param("id")

	if err := controller.db.First(&contributor, contributorId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Nie znaleziono kontrybutora"})
		return
	}

	currentStatus := false
	if contributor.Blocked != nil {
		currentStatus = *contributor.Blocked
	}

	newStatus := !currentStatus

	if err := controller.db.Model(&contributor).Update("blocked", newStatus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd zapisu"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"blocked": newStatus})
}

type ContributorPasswordForm struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (controller GlobalController) UpdateContributorPassword(ctx *gin.Context) {
	contributorId := ctx.Param("id")
	var passwordForm ContributorPasswordForm

	if err := ctx.ShouldBindJSON(&passwordForm); err != nil {
		ctx.JSON(400, gin.H{"error": "Nieprawidłowe dane: " + err.Error()})
		return
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(passwordForm.Password), 12)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	hashedPassword := string(bytes)
	passwordForm.Password = hashedPassword

	result := controller.db.Model(&models.Contributor{}).Where("id = ?", contributorId).Updates(passwordForm)

	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": "Błąd podczas aktualizacji bazy"})
		return
	}

	if result.RowsAffected == 0 {
		ctx.JSON(404, gin.H{"error": "Nie znaleziono quizu o podanym ID"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Hasło zostało zaaktualizowane", "id": contributorId})

}

type ContributorDetails struct {
	Name        string `json:"name" binding:"required,min=3"`
	Description string `json:"description"`
	Login       string `json:"login"`
	Permissions string `json:"permissions"`
}

func (controller GlobalController) UpdateContributorDetails(ctx *gin.Context) {
	contributorId := ctx.Param("id")
	var contributor ContributorDetails

	if err := ctx.ShouldBindJSON(&contributor); err != nil {
		ctx.JSON(400, gin.H{"error": "Nieprawidłowe dane: " + err.Error()})
		return
	}

	result := controller.db.Model(&models.Contributor{}).Where("id = ?", contributorId).Updates(contributor)

	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": "Błąd podczas aktualizacji bazy"})
		return
	}

	if result.RowsAffected == 0 {
		ctx.JSON(404, gin.H{"error": "Nie znaleziono quizu o podanym ID"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Informacje zostały zaktualizowane", "id": contributorId})
}

func (controller GlobalController) DeleteContributor(ctx *gin.Context) {
	id := ctx.Param("id")
	var contributor models.Contributor

	if err := controller.db.First(&contributor, id).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Contributor not found"})
		return
	}

	err := controller.db.Select("Roles").Delete(&contributor).Error

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to delete"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Contributor and their roles associations removed"})
}

func (controller GlobalController) UpdateContributorPermissions(ctx *gin.Context) {
	contributorId := ctx.Param("id")
	var permissions struct {
		Permissions string `json:"permissions"`
	}

	if err := ctx.ShouldBindJSON(&permissions); err != nil {
		ctx.JSON(400, gin.H{"error": "Nieprawidłowe dane: " + err.Error()})
		return
	}

	result := controller.db.Model(&models.Contributor{}).Where("id = ?", contributorId).Updates(permissions)

	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": "Błąd podczas aktualizacji bazy"})
		return
	}

	if result.RowsAffected == 0 {
		ctx.JSON(404, gin.H{"error": "Nie znaleziono quizu o podanym ID"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Informacje zostały zaktualizowane", "id": contributorId})
}

func (controller GlobalController) CreateContributor(c *gin.Context) {
	var contributor models.Contributor
	if err := c.ShouldBindJSON(&contributor); err != nil {
		logger.ErrorLog(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := controller.db.Create(&contributor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Aktualizacja hasła jest pomyślna!"})
}

// post
func (controller GlobalController) CheckPassword(c *gin.Context) {
	var payload struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	var contributor models.Contributor

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err, "correct": false})
		return
	}

	result := controller.db.Model(models.Contributor{}).Where("login = ?", payload.Login).Find(&contributor)
	if result.Error != nil {
		logger.ErrorLog("Incorrect password")
		c.JSON(400, gin.H{"message": result.Error, "correct": false})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(contributor.Password), []byte(payload.Password))
	if err != nil {
		logger.ErrorLog("Incorrect password")
		c.JSON(400, gin.H{"message": "Niepoprawne hasło", "correct": false})
		return
	}

	c.JSON(200, gin.H{"message": "Poprawne hasło!", "correct": true})
}
