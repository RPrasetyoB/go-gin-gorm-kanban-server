package models

import "gorm.io/gorm"

type Items struct {
	gorm.Model
	Todo_id             int    `gorm:"not null"`
	Name                string `gorm:"not null" json:"name"`
	Progress_percentage int    `gorm:"not null" json:"progress_percentage"`
}

func (Items) TableName() string {
	return "Items"
}
