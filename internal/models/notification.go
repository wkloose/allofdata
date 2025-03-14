package models

import (
    "gorm.io/gorm"
    "time"
)

type Notification struct {
    gorm.Model
    UserID  uint      `gorm:"not null"` // Foreign Key ke User
    User    User      `gorm:"foreignKey:UserID"`
    Title   string    `gorm:"not null"` 
    Message string    `gorm:"not null"` 
    Read    bool      `gorm:"default:false"` // Status apakah notifikasi sudah dibaca
    Time    time.Time `gorm:"not null"` 
}
