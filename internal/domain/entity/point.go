package entity

import "gorm.io/gorm"

type Point struct {
    gorm.Model
    UserID uint   `gorm:"not null"` // Foreign Key ke User
    User   User   `gorm:"foreignKey:UserID"`
    Points int    `gorm:"not null;default:0"` 
}
