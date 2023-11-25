package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Id uint `gorm:"primaryKey, autoIncrement"` 
	Email string `gorm:"unique, size:255"`
	Password string `gorm:"not null, size:50"`
	Role string `gorm:"default:user, size:10"`
	IsAuth bool `gorm:"default:false"`
}



