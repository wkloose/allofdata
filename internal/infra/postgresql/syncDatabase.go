package postgresql

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string	`gorm:"unique"` 
	Password string
}
func SyncDatabase() {
	DB.AutoMigrate(&User{})
}
