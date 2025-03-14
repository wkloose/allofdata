package postgresql

import (
	"trashure/internal/models"

	"gorm.io/gorm"
)
func SyncDatabase(db *gorm.DB) error {
	if err:= db.AutoMigrate(
		&models.User{},
		&models.UserHistory{},
		&models.TrashureRequest{},
		&models.Waste{},
		&models.WasteCollection{},
		&models.Education{},
		&models.Transaction{},
		&models.GreenActivity{},
		&models.Point{},
		&models.Notification{},
		&models.Quiz{},
	);err != nil{
		return err
	}
	return nil
}
