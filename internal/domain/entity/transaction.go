package entity

import (
	"time"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID     uint        `gorm:"not null"`
	User       User        `gorm:"foreignKey:UserID"`
	WasteID    uint        `gorm:"not null"`
	Waste      Waste       `gorm:"foreignKey:WasteID"`
	Weight     float64     `gorm:"not null"` // Berat sampah dalam kg
	TotalPrice float64     `gorm:"not null"`
	Date       time.Time   `gorm:"not null"`
	Status     string      `gorm:"default:pending"` // "pending", "confirmed", "rejected"
}
