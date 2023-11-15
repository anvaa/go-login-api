package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Id uint `gorm:"primaryKey, autoIncrement"` 
	Email string `gorm:"unique, size:255"`
	Password string `gorm:"not null, size:100"`
	IsAuth bool `gorm:"default:false"`
	IsAdmin bool `gorm:"default:false"`
}



