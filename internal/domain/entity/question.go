package entity

import (
    "gorm.io/gorm"
)

type Question struct {
    gorm.Model
    QuizID      uint   `gorm:"not null"` // ID Kuis (relasi ke Quiz)
    Question    string `gorm:"not null"` 
    OptionA     string `gorm:"not null"` 
    OptionB     string `gorm:"not null"` 
    OptionC     string `gorm:"not null"` 
    OptionD     string `gorm:"not null"` 
    CorrectAnswer string `gorm:"not null"` 
}
