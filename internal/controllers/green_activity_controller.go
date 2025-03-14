package controllers

import (
    "net/http"
    "time"
    "trashure/internal/models"
    "trashure/internal/postgresql"

    "github.com/gin-gonic/gin"
)

// GetAllGreenActivities - Mendapatkan semua aktivitas hijau
func GetAllGreenActivities(c *gin.Context) {
    var activities []models.GreenActivity
    if err := postgresql.DB.Find(&activities).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve activities"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": activities})
}

// CreateGreenActivity - Menambahkan aktivitas hijau baru
func CreateGreenActivity(c *gin.Context) {
    var body struct {
        Title       string    `json:"title" binding:"required"`
        Description string    `json:"description"`
        Location    string    `json:"location" binding:"required"`
        Date        time.Time `json:"date" binding:"required"`
    }

    if err := c.Bind(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
        return
    }

    activity := models.GreenActivity{
        Title:       body.Title,
        Description: body.Description,
        Location:    body.Location,
        Date:        body.Date,
    }

    if err := postgresql.DB.Create(&activity).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create activity"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Green activity created successfully", "data": activity})
}

// UpdateGreenActivity - Memperbarui aktivitas hijau
func UpdateGreenActivity(c *gin.Context) {
    id := c.Param("id")
    var activity models.GreenActivity

    if err := postgresql.DB.First(&activity, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
        return
    }

    if err := c.ShouldBindJSON(&activity); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
        return
    }

    if err := postgresql.DB.Save(&activity).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update activity"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Green activity updated successfully", "data": activity})
}

// DeleteGreenActivity - Menghapus aktivitas hijau
func DeleteGreenActivity(c *gin.Context) {
    id := c.Param("id")
    if err := postgresql.DB.Delete(&models.GreenActivity{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete activity"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Green activity deleted successfully", "id": id})
}

// RecommendGreenActivities - Rekomendasi aktivitas hijau berdasarkan filter
func RecommendGreenActivities(c *gin.Context) {
    var activities []models.GreenActivity
    location := c.Query("location")
    startDate := c.Query("start_date")
    endDate := c.Query("end_date")

    query := postgresql.DB

    if location != "" {
        query = query.Where("location = ?", location)
    }
    if startDate != "" && endDate != "" {
        query = query.Where("date BETWEEN ? AND ?", startDate, endDate)
    }

    if err := query.Find(&activities).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve recommendations"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": activities})
}
