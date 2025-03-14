package models

import (
	"gorm.io/gorm"
	"time"
)

type GreenActivity struct {
	gorm.Model
	Title       string    `gorm:"not null"`
	Description string
	Location    string    `gorm:"not null"`
	Date        time.Time `gorm:"not null"`
}
