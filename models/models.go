package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Id         int    `gorm:"primaryKey, autoIncrement"`
	Email      string `gorm:"unique, size:255"`
	Password   string `gorm:"not null, size:255"`
	Role       string `gorm:"default:user, size:20"`
	IsAuth     bool   `gorm:"default:false"`
	AccessTime int    `gorm:"default:3600"` // 1 hour
}

type Links struct {
	gorm.Model
	Id     int    `gorm:"primaryKey, autoIncrement"`
	Url    string `gorm:"default:/v/appstart"`
	UserId int    `gorm:"default:0"`
}
