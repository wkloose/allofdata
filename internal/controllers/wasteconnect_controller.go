package controllers

import (
	"net/http"
	"trashure/internal/models"
	"trashure/internal/postgresql"

	"github.com/gin-gonic/gin"
)
func ConfirmOrder(c *gin.Context) {
    id := c.Param("id")
    var body struct {
        Status string `json:"status" binding:"required"` // "confirmed" atau "rejected"
    }

    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
        return
    }

    var request models.TrashureRequest
    if err := postgresql.DB.First(&request, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Trashure Request not found"})
        return
    }

    // Update status permintaan
    request.Status = body.Status
    if err := postgresql.DB.Save(&request).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Trashure Request"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Trashure Request status updated successfully", "data": request})
}
func RateOrder(c *gin.Context) {
    id := c.Param("id")
    var body struct {
        Points int `json:"points" binding:"required"` // Jumlah poin yang diberikan
    }

    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
        return
    }

    var request models.WasteConnect
    if err := postgresql.DB.First(&request, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Trashure Request not found"})
        return
    }

    // Update poin
    request.Points = body.Points
    if err := postgresql.DB.Save(&request).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update points for Trashure Request"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Points added successfully", "data": request})
}

