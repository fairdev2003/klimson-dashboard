package services

import "github.com/zgierz/harc_quiz/backend/models"

type QuizService interface {
	GetAdminQuizes() ([]models.Quiz, error)
	GetPublicQuizes() ([]models.PublicQuiz, error)
	CreateQuiz(*models.Quiz) (*models.Quiz, error)
	AddMultipleQuestions(quizID uint, questions []models.Question) (*models.Quiz, error)
	AddQuestion(quizID uint, question models.Question) (*models.Quiz, error)
	CheckAnswer(questionId uint, answerId uint) (bool, error)
}
