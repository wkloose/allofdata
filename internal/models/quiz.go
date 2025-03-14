package models

import (
    "gorm.io/gorm"
)

type Quiz struct {
    gorm.Model
    Title       string     `gorm:"not null"`
    Description string     `gorm:"not null"` 
    Questions   []Question `gorm:"foreignKey:QuizID"` // Relasi ke tabel Question
}
