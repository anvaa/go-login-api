package initializers

import (
	"models"
)

func SyncDB() {
	DB.AutoMigrate(
		&models.Users{},
	
	)
}