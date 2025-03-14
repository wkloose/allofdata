package postgresql

import (
	"trashure/internal/domain/entity"
)
func SyncDatabase() {
	DB.AutoMigrate(&entity.User{})
	DB.AutoMigrate(&entity.UserHistory{})
	DB.AutoMigrate(&entity.TrashureRequest{})
	DB.AutoMigrate(&entity.Waste{})
	DB.AutoMigrate(&entity.WasteCollection{})
	DB.AutoMigrate(&entity.Education{})
	DB.AutoMigrate(&entity.Transaction{})
	DB.AutoMigrate(&entity.GreenActivity{})
	DB.AutoMigrate(&entity.Point{})
	DB.AutoMigrate(&entity.Notification{})
	DB.AutoMigrate(&entity.Quiz{})
}
