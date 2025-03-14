package controllers

import (
    "net/http"
    "trashure/internal/domain/entity"
    "trashure/internal/infra/postgresql"

    "github.com/gin-gonic/gin"
)

// GetPoints - Melihat poin pengguna
func GetPoints(c *gin.Context) {
    var points []entity.Point
    if err := postgresql.DB.Preload("User").Find(&points).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve points"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": points})
}

// AddPoints - Menambahkan poin ke pengguna
func AddPoints(c *gin.Context) {
    var body struct {
        UserID uint `json:"user_id" binding:"required"`
        Points int  `json:"points" binding:"required"`
    }

    if err := c.Bind(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
        return
    }

    var point entity.Point
    postgresql.DB.Where("user_id = ?", body.UserID).First(&point)

    // Jika pengguna tidak memiliki catatan poin, buat catatan baru
    if point.ID == 0 {
        point = entity.Point{
            UserID: body.UserID,
            Points: body.Points,
        }
        postgresql.DB.Create(&point)
    } else {
        // Tambahkan poin ke catatan yang ada
        point.Points += body.Points
        postgresql.DB.Save(&point)
    }

    c.JSON(http.StatusOK, gin.H{"message": "Points added successfully", "data": point})
}

// RedeemPoints - Menukarkan poin
func RedeemPoints(c *gin.Context) {
    var body struct {
        UserID uint `json:"user_id" binding:"required"`
        Points int  `json:"points" binding:"required"`
    }

    if err := c.Bind(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
        return
    }

    var point entity.Point
    postgresql.DB.Where("user_id = ?", body.UserID).First(&point)

    if point.ID == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User does not have points"})
        return
    }

    // Cek apakah poin mencukupi
    if point.Points < body.Points {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough points"})
        return
    }

    // Kurangi poin
    point.Points -= body.Points
    postgresql.DB.Save(&point)

    c.JSON(http.StatusOK, gin.H{"message": "Points redeemed successfully", "data": point})
}
