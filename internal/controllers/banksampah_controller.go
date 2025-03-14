package controllers

import (
    "net/http"
    "trashure/internal/models"
    "trashure/internal/postgresql"

    "github.com/gin-gonic/gin"
)


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

    bankSampah := models.BankSampah{
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

func GetAllBankSampah(c *gin.Context) {
    var bankSampahList []models.BankSampah
    if err := postgresql.DB.Find(&bankSampahList).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve Bank Sampah data"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": bankSampahList})
}

func GetBankSampahByID(c *gin.Context) {
    id := c.Param("id")
    var bankSampah models.BankSampah

    if err := postgresql.DB.First(&bankSampah, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Bank Sampah not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": bankSampah})
}


func UpdateBankSampahByID(c *gin.Context) {
    id := c.Param("id")
    var bankSampah models.BankSampah

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

func DeleteBankSampahByID(c *gin.Context) {
    id := c.Param("id")
    var bankSampah models.BankSampah

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


func GetBankSampahCollections(c *gin.Context) {
    user, _ := c.Get("user")
    currentUser := user.(models.User)

    var collections []models.WasteCollection
    if err := postgresql.DB.Where("bank_sampah_id = ?", currentUser.ID).Find(&collections).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve collections"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": collections})
}

func UpdateCollectionStatusByBankSampah(c *gin.Context) {
    id := c.Param("id")
    var body struct {
        Status string `json:"status" binding:"required"`
    }

    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
        return
    }

    var collection models.WasteCollection
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

func GetTrashureRequestsByBankSampah(c *gin.Context) {
    var requests []models.TrashureRequest

    if err := postgresql.DB.Where("status = ?", "pending").Find(&requests).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve trashure requests"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": requests})
}
