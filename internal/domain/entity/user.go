package entity

import "gorm.io/gorm"
type User struct {
	gorm.Model
	Name  string `gorm:"not null"` 
 	Email  string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Province    string
    City        string
    District    string
    SubDistrict string
    Address     string
    Points      int `gorm:"default:0"` 
	DateOfBirth  string `gorm:"not null"`      // Tanggal lahir (format YYYY-MM-DD)        
    BankAccount  string `gorm:"unique;not null"`
	Role        string `gorm:"not null;default:user"`
}