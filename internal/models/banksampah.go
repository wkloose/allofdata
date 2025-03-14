package models

import (
    "gorm.io/gorm"
)

type BankSampah struct {
    gorm.Model
    Title           string `gorm:"not null"` 
    Location        string `gorm:"not null"` 
    PickupTime      string `gorm:"not null"` // Jam penjemputan (contoh: "08:00 - 16:00")
    PickupDays      string `gorm:"not null"` // Hari penjemputan (contoh: "Senin, Rabu, Jumat")
}
