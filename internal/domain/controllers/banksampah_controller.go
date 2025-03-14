package controllers

import (
    "net/http"
    "trashure/internal/domain/entity"
    "trashure/internal/infra/postgresql"

    "github.com/gin-gonic/gin"
)

// CreateBankSampah - Membuat Bank Sampah baru
func CreateBankSampah(c *gin.Context) {
    var body struct {
        Title      string `json:"title" binding:"required"`
        Location   string `json:"location" binding:"required"`
        PickupTime string `json:"pickup_time" binding:"required"`
        PickupDays string `json:"pickup_days" binding:"required"`
    }

    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
        return
    }

    bankSampah := entity.BankSampah{
        Title:      body.Title,
        Location:   body.Location,
        PickupTime: body.PickupTime,
        PickupDays: body.PickupDays,
    }

    if err := postgresql.DB.Create(&bankSampah).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Bank Sampah"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Bank Sampah created successfully", "data": bankSampah})
}

// GetAllBankSampah - Mendapatkan semua Bank Sampah
func GetAllBankSampah(c *gin.Context) {
    var bankSampahList []entity.BankSampah
    if err := postgresql.DB.Find(&bankSampahList).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve Bank Sampah data"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": bankSampahList})
}

// GetBankSampahByID - Mendapatkan detail Bank Sampah berdasarkan ID
func GetBankSampahByID(c *gin.Context) {
    id := c.Param("id")
    var bankSampah entity.BankSampah

    if err := postgresql.DB.First(&bankSampah, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Bank Sampah not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": bankSampah})
}

// UpdateBankSampahByID - Memperbarui Bank Sampah berdasarkan ID
func UpdateBankSampahByID(c *gin.Context) {
    id := c.Param("id")
    var bankSampah entity.BankSampah

    if err := postgresql.DB.First(&bankSampah, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Bank Sampah not found"})
        return
    }

    var body struct {
        Title      string `json:"title"`
        Location   string `json:"location"`
        PickupTime string `json:"pickup_time"`
        PickupDays string `json:"pickup_days"`
    }

    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
        return
    }

    if body.Title != "" {
        bankSampah.Title = body.Title
    }
    if body.Location != "" {
        bankSampah.Location = body.Location
    }
    if body.PickupTime != "" {
        bankSampah.PickupTime = body.PickupTime
    }
    if body.PickupDays != "" {
        bankSampah.PickupDays = body.PickupDays
    }

    if err := postgresql.DB.Save(&bankSampah).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Bank Sampah"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Bank Sampah updated successfully", "data": bankSampah})
}

// DeleteBankSampahByID - Menghapus Bank Sampah berdasarkan ID
func DeleteBankSampahByID(c *gin.Context) {
    id := c.Param("id")
    var bankSampah entity.BankSampah

    if err := postgresql.DB.First(&bankSampah, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Bank Sampah not found"})
        return
    }

    if err := postgresql.DB.Delete(&bankSampah).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Bank Sampah"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Bank Sampah deleted successfully"})
}


// GetBankSampahCollections - Mendapatkan semua jadwal penjemputan untuk Bank Sampah
func GetBankSampahCollections(c *gin.Context) {
    user, _ := c.Get("user")
    currentUser := user.(entity.User)

    var collections []entity.WasteCollection
    if err := postgresql.DB.Where("bank_sampah_id = ?", currentUser.ID).Find(&collections).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve collections"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": collections})
}

// UpdateCollectionStatusByBankSampah - Mengubah status penjemputan (done, pending, canceled)
func UpdateCollectionStatusByBankSampah(c *gin.Context) {
    id := c.Param("id")
    var body struct {
        Status string `json:"status" binding:"required"` // Status baru: done, pending, canceled
    }

    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
        return
    }

    var collection entity.WasteCollection
    if err := postgresql.DB.First(&collection, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Collection not found"})
        return
    }

    collection.Status = body.Status
    if err := postgresql.DB.Save(&collection).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update collection status"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Collection status updated successfully", "data": collection})
}

// GetTrashureRequestsByBankSampah - Melihat semua Trashure Requests terkait Bank Sampah
func GetTrashureRequestsByBankSampah(c *gin.Context) {
    var requests []entity.TrashureRequest

    if err := postgresql.DB.Where("status = ?", "pending").Find(&requests).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve trashure requests"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": requests})
}
