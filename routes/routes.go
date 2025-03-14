package routes

import (
    "trashure/internal/domain/controllers"
    "trashure/internal/framework/middleware"

    "github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
    // Public Routes (tanpa autentikasi)
    auth := r.Group("/")
    {
        auth.POST("/signup", controllers.Signup)               // Registrasi pengguna
        auth.POST("/login", controllers.Login)                 // Login pengguna
        auth.GET("/validate", middleware.RequireAuth, controllers.Validate) // Validasi token pengguna
    }

    // User Routes
    users := r.Group("/users")
    users.Use(middleware.RequireAuth) // Hanya pengguna yang terautentikasi
    {
        users.POST("/", middleware.AdminOnly, controllers.CreateUser)           // Admin membuat pengguna baru
        users.GET("/", middleware.AdminOnly, controllers.GetAllUsers)           // Admin melihat semua pengguna
        users.GET("/:id", controllers.GetUserByID)                              // Mendapatkan detail pengguna
        users.PUT("/:id", controllers.UpdateUserByID)                           // Memperbarui data pengguna
        users.DELETE("/:id", middleware.AdminOnly, controllers.DeleteUserByID)  // Admin menghapus pengguna
        users.GET("/ranking", controllers.GetUserRanking)                       // Mendapatkan peringkat pengguna berdasarkan poin
        users.GET("/:id/history", controllers.GetUserHistory)                   // Mendapatkan riwayat aktivitas pengguna
        users.POST("/:id/points", controllers.AddPoints)                        // Menambahkan poin pengguna
    }

    // Trashure Request Routes
    trashure := r.Group("/trashure-requests")
    trashure.Use(middleware.RequireAuth) // Mengamankan rute dengan autentikasi
    {
        trashure.POST("/", controllers.CreateTrashureRequest)                   // User membuat permintaan
        trashure.GET("/", middleware.AdminOnly, controllers.GetTrashureRequest) // Admin melihat semua permintaan
        trashure.PUT("/:id/confirm", middleware.AdminOnly, controllers.ConfirmTrashureRequest) // Admin mengkonfirmasi permintaan
        trashure.GET("/:id", controllers.GetTrashureRequest)                    // User mendapatkan detail permintaan
        trashure.DELETE("/:id", controllers.DeleteTrashureRequest)              // Admin atau User menghapus permintaan
        trashure.GET("/signed-url", controllers.GetSignedURL)                   // Mendapatkan Signed URL untuk file
    }

    // EduGreen Routes
    edugreen := r.Group("/edugreen")
    edugreen.Use(middleware.RequireAuth) // Semua rute EduGreen dilindungi autentikasi
    {
        // Video Routes
        edugreen.GET("/videos", controllers.GetVideos)                   // User melihat semua video edukasi
        edugreen.POST("/videos", middleware.AdminOnly, controllers.CreateVideo) // Admin menambah video edukasi
        edugreen.PUT("/videos/:id", middleware.AdminOnly, controllers.UpdateVideo) // Admin memperbarui video
        edugreen.DELETE("/videos/:id", middleware.AdminOnly, controllers.DeleteVideo) // Admin menghapus video
        edugreen.POST("/videos/:id/complete", controllers.CompleteVideo) // User menyelesaikan video edukasi

        // Quiz Routes
        edugreen.GET("/quizzes", controllers.GetQuizzes)                        // User melihat daftar kuis
        edugreen.POST("/quizzes", middleware.AdminOnly, controllers.CreateQuiz) // Admin membuat kuis
        edugreen.POST("/quizzes/:id/complete", controllers.CompleteQuiz)     
    }

    // Bank Sampah Routes
    banksampah := r.Group("/banksampah")
    banksampah.Use(middleware.RequireAuth, middleware.BankSampahOnly) // Autentikasi khusus untuk bank sampah
    {
        banksampah.GET("/collections", controllers.GetBankSampahCollections)               // Mendapatkan jadwal penjemputan terkait bank sampah
        banksampah.PUT("/collections/:id/status", controllers.UpdateCollectionStatusByBankSampah) // Mengubah status penjemputan
        banksampah.GET("/trashure-requests", controllers.GetTrashureRequestsByBankSampah)  // Melihat Trashure Requests yang terkait
    }

    // Admin Routes
    admin := r.Group("/admin")
    admin.Use(middleware.RequireAuth, middleware.AdminOnly) // Rute untuk admin
    {
        admin.GET("/statistics", controllers.GetStatistics) // Admin melihat laporan statistik aplikasi
    }
}
