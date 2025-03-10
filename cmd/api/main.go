package main

import (
		"fmt"
	"trashure/internal/domain/user"
	"trashure/internal/framework"
	"trashure/internal/infra/postgresql"
	"os"
	"github.com/gin-gonic/gin"
)

func init() {
	postgresql.ConnectToDb()
	postgresql.SyncDatabase()
}
func main() {
	r := gin.Default()
	r.POST("/signup", user.Signup)
	r.POST("/login", user.Login)
	r.GET("/validate", framework.RequireAuth, user.Validate)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3008"
	}
	
	r.Run(fmt.Sprintf(":%s", port))
}
