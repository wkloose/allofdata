package routes

import (
    "trashure/internal/controllers"
    "trashure/internal/middleware"

    "github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
    auth := r.Group("/")
    {
        auth.POST("/signup", controllers.Signup)               // Registrasi pengguna
        auth.POST("/login", controllers.Login)                 // Login pengguna
        auth.GET("/validate", middleware.RequireAuth, controllers.Validate) // Validasi token pengguna
    }

    // User Routes
    users := r.Group("/users")
    users.Use(middleware.RequireAuth) // Perlu TOKEN
    {
        users.POST("/", middleware.AdminOnly, controllers.CreateUser)           // AdminOnly Bisa membuat pengguna baru
        users.GET("/", middleware.AdminOnly, controllers.GetAllUsers)           // AdminOnly Melihat semua pengguna
        users.GET("/:id", controllers.GetUserByID)                              // Mendapatkan Detail pengguna berdasarkan id
        users.PUT("/:id", controllers.UpdateUserByID)                           // Memperbarui Data pengguna berdasarkan id
        users.DELETE("/:id", middleware.AdminOnly, controllers.DeleteUserByID)  // Admin Menghapus pengguna
        users.GET("/ranking", controllers.GetUserRanking)                       // Menggunakan Descending untuk Menyusun data
        users.GET("/:id/history", controllers.GetUserHistory)                   // Mendapatkan riwayat aktivitas pengguna
        users.POST("/:id/points", controllers.AddPoints)                        // Menambahkan Otomatis Poin pengguna
    }

    trashure := r.Group("/trashure-requests")
    trashure.Use(middleware.RequireAuth) 
    {
        trashure.POST("/", controllers.CreateTrashureRequest)                   
        trashure.GET("/", middleware.AdminOnly, controllers.GetTrashureRequest) 
        trashure.PUT("/:id/confirm", middleware.AdminOnly, controllers.ConfirmTrashureRequest) 
        trashure.GET("/:id", controllers.GetTrashureRequest)                    
        trashure.DELETE("/:id", controllers.DeleteTrashureRequest)             
        trashure.GET("/signed-url", controllers.GetSignedURL)                  
    }

    
    edugreen := r.Group("/edugreen")
    edugreen.Use(middleware.RequireAuth) 
    {
        edugreen.GET("/videos", controllers.GetVideos)                   
        edugreen.POST("/videos", middleware.AdminOnly, controllers.CreateVideo) 
        edugreen.PUT("/videos/:id", middleware.AdminOnly, controllers.UpdateVideo) 
        edugreen.DELETE("/videos/:id", middleware.AdminOnly, controllers.DeleteVideo) 
        edugreen.POST("/videos/:id/complete", controllers.CompleteVideo) 

        edugreen.GET("/quizzes", controllers.GetQuizzes)                       
        edugreen.POST("/quizzes", middleware.AdminOnly, controllers.CreateQuiz) 
        edugreen.POST("/quizzes/:id/complete", controllers.CompleteQuiz)     
    }

    banksampah := r.Group("/banksampah")
    banksampah.Use(middleware.RequireAuth, middleware.BankSampahOnly)
    {
        banksampah.GET("/collections", controllers.GetBankSampahCollections)              
        banksampah.PUT("/collections/:id/status", controllers.UpdateCollectionStatusByBankSampah)
        banksampah.GET("/trashure-requests", controllers.GetTrashureRequestsByBankSampah)  // Melihat Trashure Requests yang terkait
    }

    admin := r.Group("/admin")
    admin.Use(middleware.RequireAuth, middleware.AdminOnly) 
    {
        admin.GET("/statistics", controllers.GetStatistics)
    }

    wasteconnect := r.Group("/wasteconnect")
    wasteconnect.Use(middleware.RequireAuth)
    {
    wasteconnect.PUT("/order/:id/confirm", middleware.BankSampahOnly, controllers.ConfirmOrder)
    wasteconnect.POST("/order/:id/rate", middleware.BankSampahOnly, controllers.RateOrder)
    wasteconnect.GET("/:id/history", controllers.GetUserHistory)
    }
}