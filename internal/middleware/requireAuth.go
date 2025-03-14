package middleware

import (
	"trashure/internal/models"
	"trashure/internal/postgresql"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
    tokenString, err := c.Cookie("Authorization")
    if err != nil {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("SECRET")), nil
    })

    if err != nil || !token.Valid {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }

    // Cari pengguna berdasarkan ID di token
    var user models.User
    postgresql.DB.First(&user, claims["sub"])

    if user.ID == 0 {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }

    c.Set("user", user)
    c.Next()
}
