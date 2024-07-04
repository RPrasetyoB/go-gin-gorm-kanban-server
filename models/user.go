package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `json:"password"`
}

func (Users) TableName() string {
	return "Users"
}
