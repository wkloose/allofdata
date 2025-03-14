package entity

import (
	"gorm.io/gorm"
	"time"
)

type Waste struct {
	gorm.Model
	Name        string  `gorm:"not null;unique"`
	Category    string  `gorm:"not null"`  // "plastik", "logam", "kertas", dll.
	PricePerKg  float64 `gorm:"not null"`  // Harga estimasi per kg
	Description string
}

type WasteCollection struct {
	gorm.Model
	UserID       uint      `gorm:"not null"`    // Foreign Key ke User
    User         User      `gorm:"foreignKey:UserID"`
    Location     string    `gorm:"not null"`    
    PickupDate   time.Time `gorm:"not null"`    
    Day          string    `gorm:"not null"`    
    BankSampahID uint      `gorm:"not null"`    
    Status       string    `gorm:"default:pending"` // Status: pending, done
	SortingMode string    `gorm:"not null"`        // "normal" atau "dipilah di tempat"
}
