package controllers

import (
    "net/http"
    "time"
    "trashure/internal/domain/entity"
    "trashure/internal/infra/postgresql"
    "github.com/gin-gonic/gin"
)

// CreateWasteCollection - Membuat jadwal penjemputan baru
func CreateWasteCollection(c *gin.Context) {
    var body struct {
        PickupDate  time.Time `json:"pickup_date" binding:"required"`
        SortingMode string    `json:"sorting_mode" binding:"required"`
    }

    if err := c.Bind(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
        return
    }

    user, _ := c.Get("user")
    currentUser := user.(entity.User)

    collection := entity.WasteCollection{
        UserID:     currentUser.ID,
        PickupDate: body.PickupDate,
        Status:     "pending",
        SortingMode: body.SortingMode,
    }

    if err := postgresql.DB.Create(&collection).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create waste collection"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Waste collection created successfully", "data": collection})
}

// GetWasteCollections - Mendapatkan semua jadwal penjemputan
func GetWasteCollections(c *gin.Context) {
    user, _ := c.Get("user")
    currentUser := user.(entity.User)

    var collections []entity.WasteCollection
    if err := postgresql.DB.Where("user_id = ?", currentUser.ID).Find(&collections).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve waste collections"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": collections})
}

// UpdateWasteCollectionStatus - Perbarui status jadwal
func UpdateWasteCollectionStatus(c *gin.Context) {
    id := c.Param("id")
    var body struct {
        Status string `json:"status" binding:"required"` // Status: completed, canceled
    }

    if err := c.Bind(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
        return
    }

    var collection entity.WasteCollection
    if err := postgresql.DB.First(&collection, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Waste collection not found"})
        return
    }

    collection.Status = body.Status
    if err := postgresql.DB.Save(&collection).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update status"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Status updated successfully", "data": collection})
}

// DeleteWasteCollection - Hapus jadwal penjemputan
func DeleteWasteCollection(c *gin.Context) {
    id := c.Param("id")

    if err := postgresql.DB.Delete(&entity.WasteCollection{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete waste collection"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Waste collection deleted successfully", "id": id})
}
