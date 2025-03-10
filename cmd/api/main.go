package main

import (
	//	"fmt"
	"trashure/internal/domain/user"
	"trashure/internal/framework"
	"trashure/internal/infra/config"
	"trashure/internal/infra/postgresql"

	//"trashure/internal/app/bootstrap"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	postgresql.ConnectToDb()
	postgresql.SyncDatabase()
}
func main() {
	r := gin.Default()
	r.POST("/signup", user.Signup)
	r.POST("/login", user.Login)
	r.GET("/validate", framework.RequireAuth, user.Validate)
	r.Run() // listen and serve on 0.0.0.0:8080
}
