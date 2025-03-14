package controllers

import (
    "net/http"
    "trashure/internal/models"
    "trashure/internal/postgresql"
    "github.com/gin-gonic/gin"
)

func CreateVideo(c *gin.Context) {
    var body struct {
        Title  string `json:"title" binding:"required"`
        Link   string `json:"link" binding:"required"`
        Points int    `json:"points" binding:"required"`
    }

    if err := c.Bind(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
        return
    }

    video := models.Education{
        Title:  body.Title,
        Link:   body.Link,
        Points: body.Points,
    }

    if err := postgresql.DB.Create(&video).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create video"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Video created successfully", "data": video})
}


func GetVideos(c *gin.Context) {
    var videos []models.Education
    if err := postgresql.DB.Find(&videos).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve videos"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": videos})
}

func CompleteVideo(c *gin.Context) {
    id := c.Param("id")
    user, _ := c.Get("user")
    currentUser := user.(models.User)

    var video models.Education
    if err := postgresql.DB.First(&video, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
        return
    }

    currentUser.Points += video.Points
    if err := postgresql.DB.Save(&currentUser).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user points"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Video completed successfully", "points_added": video.Points})
}

func UpdateVideo(c *gin.Context) {
    id := c.Param("id")
    var video models.Education

    if err := postgresql.DB.First(&video, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
        return
    }

    var body struct {
        Title  string `json:"title"`
        Link   string `json:"link"`
        Points int    `json:"points"`
    }

    if err := c.Bind(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
        return
    }

    if body.Title != "" {
        video.Title = body.Title
    }
    if body.Link != "" {
        video.Link = body.Link
    }
    if body.Points > 0 {
        video.Points = body.Points
    }

    if err := postgresql.DB.Save(&video).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update video"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Video updated successfully", "data": video})
}

func DeleteVideo(c *gin.Context) {
    id := c.Param("id")
    var video models.Education

    if err := postgresql.DB.First(&video, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
        return
    }

    if err := postgresql.DB.Delete(&video).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete video"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Video deleted successfully", "id": id})
}
