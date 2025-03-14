package main

import (
	"fmt"
	"os"
	"trashure/internal/infra/postgresql"
	"trashure/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	postgresql.ConnectToDb()
	postgresql.SyncDatabase()
	postgresql.SeedUsers() 
}
func main() {
	r := gin.Default()
	routes.RegisterRoutes(r)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3008"
	}
	
	r.Run(fmt.Sprintf(":%s", port))
}
