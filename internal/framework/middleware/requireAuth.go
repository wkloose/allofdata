package middleware

import (
	"trashure/internal/domain/entity"
	"trashure/internal/infra/postgresql"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
    // Ambil cookie token
    tokenString, err := c.Cookie("Authorization")
    if err != nil {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }

    // Parse token JWT
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("SECRET")), nil
    })

    if err != nil || !token.Valid {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }

    // Ambil klaim dari token
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }

    // Cari pengguna berdasarkan ID di token
    var user entity.User
    postgresql.DB.First(&user, claims["sub"])

    if user.ID == 0 {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }

    // Tambahkan pengguna ke konteks
    c.Set("user", user)
    c.Next()
}
