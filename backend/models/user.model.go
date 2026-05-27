package models

type User struct {
	BaseModel
	Firstname      string `json:"first_name"`
	Lastname       string `json:"last_name"`
	Nickname       string `json:"nickname"`
	Password       string `json:"password"`
	RoleID         uint   `json:"role_id"`
	Role           Role   `json:"role" gorm:"foreignKey:RoleID"`
	ProfilePicture string `json:"pfp"`
	Blocked        *bool  `json:"blocked"`
}

type Role struct {
	BaseModel
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions" gorm:"foreignKey:RoleID"`
}

type Permission struct {
	BaseModel
	Name   string `json:"name"`
	RoleID uint   `json:"role_id"`
}
