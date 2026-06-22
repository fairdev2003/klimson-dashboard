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

type MigrationType struct {
	Model any
	Name  string
	Icon  string
}

var MigratableModels = []MigrationType{
	{Model: &User{}, Name: "Users", Icon: "mdi:user"},
	{Model: &Permission{}, Name: "Permissions", Icon: "material-symbols:key-rounded"},
	{Model: &Role{}, Name: "Role", Icon: "mingcute:hat-fill"},
	{Model: &ContextStorage{}, Name: "Context Storage", Icon: "material-symbols:contextual-token"},
	{Model: &Storage{}, Name: "File Storage", Icon: "material-symbols:storage"},
	{Model: &Todo{}, Name: "Todo", Icon: "ri:todo-fill"},
	{Model: &TodoTask{}, Name: "Todo Task", Icon: "idk really"},
}
