package postgresql

import (
    "log"
    "trashure/internal/models"

    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)


func SeedUsers(db *gorm.DB) {
    var count int64
    db.Model(&models.User{}).Count(&count)
    if count > 0 {
        return
    }

    password, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

    users := []models.User{
        {
            Name:        "Admin",
            Email:       "admin@example.com",
            Password:    string(password),
            Role:        "admin",
            Province:    "DKI Jakarta",
            City:        "Jakarta Pusat",
            District:    "Gambir",
            SubDistrict: "Cideng",
            Address:     "Admin Address",
            Points:      0,
            DateOfBirth: "1990-01-01",
            BankAccount: "123456789",
        },
        {
            Name:        "Bank Sampah",
            Email:       "banksampah@example.com",
            Password:    string(password),
            Role:        "banksampah",
            Province:    "Jawa Tengah",
            City:        "Semarang",
            District:    "Tembalang",
            SubDistrict: "Sendangmulyo",
            Address:     "Bank Sampah Address",
            Points:      0,
            DateOfBirth: "1985-01-01",
            BankAccount: "987654321",
        },
        {
            Name:        "User",
            Email:       "user@example.com",
            Password:    string(password),
            Role:        "user",
            Province:    "Jawa Timur",
            City:        "Malang",
            District:    "Klojen",
            SubDistrict: "Oro-Oro Dowo",
            Address:     "User Address",
            Points:      0,
            DateOfBirth: "2000-01-01",
            BankAccount: "1122334455",
        },
    }

    // loop untuk memasukkan data pengguna ke dalam database
    for _, user := range users {
        var existing models.User
        result := db.First(&existing, "email = ?", user.Email)

        if result.RowsAffected == 0 {
            if err := db.Create(&user).Error; err != nil {
                log.Printf("Gagal menambahkan pengguna %s: %v\n", user.Email, err)
            } else {
                log.Printf("Pengguna %s berhasil ditambahkan\n", user.Email)
            }
        }
    }
}
