package models

type Hero struct {
	BaseModel
	Quote  string `json:"quote"`
	Image  string `json:"image_url"`
	Author string `json:"author"`
}
