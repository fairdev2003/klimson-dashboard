package models

import "time"

type Contributor struct {
	BaseModel
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Login        string    `json:"login"`
	Password     string    `json:"password" gorm:"type:text"`
	Permissions  string    `json:"permissions"`
	Thumbnail    string    `json:"thumbnail"`
	ProfileImage string    `json:"profile_image"`
	LastLogin    time.Time `json:"last_login"`
	Blocked      *bool     `json:"blocked_till"`
	Logs         []Log     `json:"logs" gorm:"type:text"`
}

type Log struct {
	BaseModel
	Action    string    `json:"action" gorm:"type:text"`
	Timestamp time.Time `json:"timestamp"`
}
