package models

type Todo struct {
	BaseModel
	Icon        string     `json:"icon"`
	Name        string     `json:"name"`
	Desc        string     `json:"desc"`
	PercentDone string     `json:"percent_done"`
	Tasks       []TodoTask `json:"tasks" gorm:"foreignKey:TodoID"`
}

type TodoTask struct {
	BaseModel
	TodoID  uint   `json:"todo_id"`
	Content string `json:"content"`
	// Todo        Todo   `json:"-" gorm:"constraint:OnDelete:CASCADE;"`
}
