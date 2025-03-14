package main

import (
	"fmt"
	"log"
	"os"
	"trashure/internal/postgresql"
	"trashure/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
    // Memuat file .env
    err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	

    // Hubungkan ke database
    db := postgresql.ConnectToDb()

    // Sinkronisasi struktur database
    if err := postgresql.SyncDatabase(db); err != nil {
        log.Fatalf("Error syncing database: %v", err)
    }
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
