package postgresql

import (
    "trashure/internal/domain/entity"
    "golang.org/x/crypto/bcrypt"
    "log"
)


func SeedUsers() {
    password, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

    users := []entity.User{
        {
            Name:        "Admin",
            Email:       "admin@example.com",
            Password:    string(password),
            Role:        "admin",
            Address:     "Admin Address",
            DateOfBirth: "1990-01-01",
            BankAccount: "123456789",
        },
        {
            Name:        "Bank Sampah",
            Email:       "banksampah@example.com",
            Password:    string(password),
            Role:        "banksampah",
            Address:     "Bank Sampah Address",
            DateOfBirth: "1985-01-01",
            BankAccount: "987654321",
        },
        {
            Name:        "User",
            Email:       "user@example.com",
            Password:    string(password),
            Role:        "user",
            Address:     "User Address",
            DateOfBirth: "2000-01-01",
            BankAccount: "1122334455",
        },
    }

    // Masukkan data ke database
    for _, user := range users {
        var existing entity.User
        result := DB.First(&existing, "email = ?", user.Email)

        // Hanya tambahkan jika belum ada di database
        if result.RowsAffected == 0 {
            if err := DB.Create(&user).Error; err != nil {
                log.Printf("Failed to seed user %s: %v\n", user.Email, err)
            } else {
                log.Printf("Seeded user %s\n", user.Email)
            }
        }
    }
}
