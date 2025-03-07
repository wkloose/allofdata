package initializers

import (
	"trashure/internal/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
