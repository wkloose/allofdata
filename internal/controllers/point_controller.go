package controllers

import (
    "net/http"
    "trashure/internal/models"
    "trashure/internal/postgresql"

    "github.com/gin-gonic/gin"
)

func GetPoints(c *gin.Context) {
    var points []models.Point
    if err := postgresql.DB.Preload("User").Find(&points).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve points"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": points})
}

func AddPoints(c *gin.Context) {
    var body struct {
        UserID uint `json:"user_id" binding:"required"`
        Points int  `json:"points" binding:"required"`
    }

    if err := c.Bind(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
        return
    }

    var point models.Point
    postgresql.DB.Where("user_id = ?", body.UserID).First(&point)

    if point.ID == 0 {
        point = models.Point{
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


func RedeemPoints(c *gin.Context) {
    var body struct {
        UserID uint `json:"user_id" binding:"required"`
        Points int  `json:"points" binding:"required"`
    }

    if err := c.Bind(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
        return
    }

    var point models.Point
    postgresql.DB.Where("user_id = ?", body.UserID).First(&point)

    if point.ID == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User does not have points"})
        return
    }

    if point.Points < body.Points {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough points"})
        return
    }

    point.Points -= body.Points
    postgresql.DB.Save(&point)

    c.JSON(http.StatusOK, gin.H{"message": "Points redeemed successfully", "data": point})
}
