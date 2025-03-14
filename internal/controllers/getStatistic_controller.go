package controllers

import (
    "net/http"
    "trashure/internal/postgresql"

    "github.com/gin-gonic/gin"
)


func GetStatistics(c *gin.Context) {
    
    var wasteStats []struct {
        Category string
        TotalKg  float64
    }
    if err := postgresql.DB.Table("wastes").
        Select("category, SUM(price_per_kg) as total_kg").
        Group("category").
        Scan(&wasteStats).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve waste statistics"})
        return
    }

    var transactionStats []struct {
        Status    string
        Total     int
    }
    if err := postgresql.DB.Table("transactions").
        Select("status, COUNT(*) as total").
        Group("status").
        Scan(&transactionStats).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve transaction statistics"})
        return
    }

    var pointStats []struct {
        Name   string
        Points int
    }
    if err := postgresql.DB.Table("points").
        Joins("JOIN users ON points.user_id = users.id").
        Select("users.name, points.points").
        Order("points.points DESC").
        Scan(&pointStats).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve point distribution statistics"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "waste_statistics":    wasteStats,
        "transaction_statistics": transactionStats,
        "point_distribution":  pointStats,
    })
}
