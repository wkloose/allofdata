package controllers

import (
	"net/http"
	"trashure/internal/models"
	"github.com/gin-gonic/gin"
)

// GetAllWaste - Ambil semua data sampah
func GetAllWaste(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List of waste"})
}

// CreateWaste - Tambah data sampah baru
func CreateWaste(c *gin.Context) {
	var newWaste models.Waste
	if err := c.ShouldBindJSON(&newWaste); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Waste created", "waste": newWaste})
}

// UpdateWaste - Perbarui data sampah
func UpdateWaste(c *gin.Context) {
	id := c.Param("id")
	var updatedWaste models.Waste
	if err := c.ShouldBindJSON(&updatedWaste); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Waste updated", "id": id, "waste": updatedWaste})
}

// DeleteWaste - Hapus data sampah
func DeleteWaste(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Waste deleted", "id": id})
}
