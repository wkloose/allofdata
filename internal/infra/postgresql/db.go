package postgresql

import "gorm.io/gorm"

// DB di-export agar bisa diakses oleh package lain
var DB *gorm.DB
