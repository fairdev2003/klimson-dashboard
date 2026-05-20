package models

type ContextStorage struct {
	BaseModel
	Key          string `json:"key"`
	Value        string `json:"value"`
	CategoryName string `json:"category_name"`
	Type         string `json:"type"`
	Public       *bool  `json:"is_public"`
	Icon         string `json:"icon"`
}
