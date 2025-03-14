package middleware

import (
    "net/http"
    "trashure/internal/domain/entity"

    "github.com/gin-gonic/gin"
)


func AdminOnly(c *gin.Context) {
    user, exists := c.Get("user")
    if !exists || user.(entity.User).Role != "admin" {
        c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: Admins only"})
        c.Abort()
        return
    }
    c.Next()
}

func BankSampahOnly(c *gin.Context) {
    user, exists := c.Get("user")
    if !exists || user.(entity.User).Role != "banksampah" {
        c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: Banksampah only"})
        c.Abort()
        return
    }
    c.Next()
}

func UserOnly(c *gin.Context) {
    user, exists := c.Get("user")
    if !exists || user.(entity.User).Role != "user" {
        c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: Users only"})
        c.Abort()
        return
    }
    c.Next()
}
