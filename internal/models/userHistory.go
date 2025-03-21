package models
import (
	 "gorm.io/gorm"
	 "time"
)
type UserHistory struct {
    gorm.Model
    UserID  uint      `gorm:"not null"` // Foreign Key ke User
    User    User      `gorm:"foreignKey:UserID"`
    Province    string
    City        string
    District    string
    SubDistrict string
    Address string    `gorm:"not null"`
    Time    time.Time `gorm:"not null"`
    Day     string    `gorm:"not null"`
}
