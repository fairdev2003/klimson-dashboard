package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Quiz struct {
	BaseModel
	Title           string        `json:"title"`
	Description     string        `json:"description" gorm:"type:text"`
	ImageURL        string        `json:"image_url"`
	EditLink        string        `json:"edit_link"`
	HasTimeLimit    *bool         `json:"has_time_limit"`
	TimeLimit       time.Duration `json:"time_limit"`
	Public          *bool         `json:"public" gorm:"default:true"`
	Difficulty      string        `json:"difficulty" gorm:"type:varchar(50)"`
	ExpectedTimeMin string        `json:"expected_time_min" gorm:"type:varchar(10)"`
	Author          string        `json:"author"`
	CompletedCount  int           `json:"completed_count" gorm:"default:0"`
	Badges          string        `json:"badges" gorm:"type:text"`
	Questions       []Question    `json:"questions" gorm:"constraint:OnDelete:CASCADE"`
	Stats           []Stat        `json:"stats"`
	Tags            string        `json:"tags" gorm:"type:text"` // tag1,tag2,tag3
}

type Question struct {
	BaseModel
	QuizID    uint     `json:"quiz_id"`
	Content   string   `json:"content" gorm:"not null"`
	ImageURL  string   `json:"image_url"`
	Type      string   `json:"type" gorm:"type:varchar(20);not null"`
	TimeLimit *int     `json:"time_limit" gorm:"column:time_limit"`
	Answers   []Answer `json:"answers" gorm:"foreignKey:QuestionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type PublicAnswer struct {
	ID         uint   `json:"id"`
	QuestionID uint   `json:"question_id"`
	Content    string `json:"content"`
}

type PublicQuestion struct {
	ID       uint           `json:"id"`
	QuizID   uint           `json:"quiz_id"`
	Content  string         `json:"content"`
	ImageURL string         `json:"image_url"`
	Type     string         `json:"type"`
	Answers  []PublicAnswer `json:"answers"`
}

type PublicQuiz struct {
	ID          uint             `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Questions   []PublicQuestion `json:"questions"`
}

type Answer struct {
	BaseModel
	QuestionID uint   `json:"question_id"`
	Content    string `gorm:"not null" json:"content"`
	IsCorrect  *bool  `gorm:"default:false" json:"is_correct,omitempty"`
}
