package models
import "gorm.io/gorm"
type WasteConnect struct {
    gorm.Model
    UserID       uint    `gorm:"not null"`           // ID pengguna
    BankSampahID uint    `gorm:"default:null"`       // ID Bank Sampah
    Type         string  `gorm:"not null"`          // Jenis sampah
    Weight       float64 `gorm:"not null"`          // Berat dalam kilogram
    Status       string  `gorm:"default:pending"`   // Status: "pending", "confirmed", "rejected"
    Points       int     `gorm:"default:0"`         // Poin yang diberikan
}
