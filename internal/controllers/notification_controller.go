package controllers

import (
    "net/http"
    "time"
    "trashure/internal/models"
    "trashure/internal/postgresql"

    "github.com/gin-gonic/gin"
)

func GetNotifications(c *gin.Context) {
    user, _ := c.Get("user")
    currentUser := user.(models.User)

    var notifications []models.Notification
    if err := postgresql.DB.Where("user_id = ?", currentUser.ID).Order("time desc").Find(&notifications).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve notifications"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": notifications})
}


func CreateNotification(c *gin.Context) {
    var body struct {
        UserID  uint   `json:"user_id" binding:"required"`
        Title   string `json:"title" binding:"required"`
        Message string `json:"message" binding:"required"`
    }

    if err := c.Bind(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
        return
    }

    notification := models.Notification{
        UserID:  body.UserID,
        Title:   body.Title,
        Message: body.Message,
        Read:    false,
        Time:    time.Now(),
    }

    if err := postgresql.DB.Create(&notification).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create notification"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Notification created successfully", "data": notification})
}

func MarkAsRead(c *gin.Context) {
    id := c.Param("id")
    var notification models.Notification

    if err := postgresql.DB.First(&notification, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Notification not found"})
        return
    }

    notification.Read = true
    postgresql.DB.Save(&notification)

    c.JSON(http.StatusOK, gin.H{"message": "Notification marked as read"})
}
	