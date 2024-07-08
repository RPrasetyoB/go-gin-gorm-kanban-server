package models

import "gorm.io/gorm"

type Todos struct {
	gorm.Model
	User_id     int    `gorm:"not null"`
	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"not null" json:"description"`
}

func (Todos) TableName() string {
	return "Todos"
}
