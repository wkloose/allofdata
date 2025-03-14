package entity

import (
    "gorm.io/gorm"
)

type TrashureRequest struct {
    gorm.Model
    UserID   uint    `gorm:"not null"`
    User     User    `gorm:"foreignKey:UserID"`
    Type     string  `gorm:"not null"`
    Weight   float64 `gorm:"not null"`
    ImageURL    string  `gorm:"not null"` // Menyimpan data gambar sebagai binary
    Price    float64
    Status   string  `gorm:"default:pending"`
}

