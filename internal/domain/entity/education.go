package entity

import "gorm.io/gorm"

type Education struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Link        string
	Points      int `gorm:"not null"`	
}
