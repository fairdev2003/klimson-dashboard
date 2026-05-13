package models

type Blog struct {
	BaseModel
	Title       string `json:"title"`
	HTML        string `json:"html"`
	Description string `json:"description"`
	Public      *bool  `json:"public"`
}
