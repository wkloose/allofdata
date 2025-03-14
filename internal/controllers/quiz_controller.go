package controllers

import (
    "net/http"
    "trashure/internal/models"
    "trashure/internal/postgresql"

    "github.com/gin-gonic/gin"
)

func CreateQuiz(c *gin.Context) {
    // Periksa apakah pengguna adalah admin
    user, _ := c.Get("user")
    currentUser := user.(models.User)
    if currentUser.Role != "admin" {
        c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can create quizzes"})
        return
    }

    // Ambil input kuis
    var body struct {
        Title       string `json:"title" binding:"required"`
        Description string `json:"description" binding:"required"`
        Questions   []struct {
            Question     string `json:"question" binding:"required"`
            OptionA      string `json:"option_a" binding:"required"`
            OptionB      string `json:"option_b" binding:"required"`
            OptionC      string `json:"option_c" binding:"required"`
            OptionD      string `json:"option_d" binding:"required"`
            CorrectAnswer string `json:"correct_answer" binding:"required"` // A, B, C, atau D
        } `json:"questions" binding:"required,len=10"` // Harus ada tepat 10 pertanyaan
    }

    // Validasi input
    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format or missing fields"})
        return
    }

    // Buat kuis
    quiz := models.Quiz{
        Title:       body.Title,
        Description: body.Description,
    }

    // Simpan kuis ke database
    if err := postgresql.DB.Create(&quiz).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create quiz"})
        return
    }

    // Tambahkan pertanyaan ke kuis
    for _, q := range body.Questions {
        question := models.Question{
            QuizID:       quiz.ID,
            Question:     q.Question,
            OptionA:      q.OptionA,
            OptionB:      q.OptionB,
            OptionC:      q.OptionC,
            OptionD:      q.OptionD,
            CorrectAnswer: q.CorrectAnswer,
        }

        if err := postgresql.DB.Create(&question).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create question"})
            return
        }
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Quiz created successfully", "data": quiz})
}
func GetQuizzes(c *gin.Context) {
    var quizzes []models.Quiz

    // Ambil semua kuis dari database
    if err := postgresql.DB.Preload("Questions").Find(&quizzes).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve quizzes"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": quizzes})
}
func CompleteQuiz(c *gin.Context) {
    id := c.Param("id")
    user, _ := c.Get("user")
    currentUser := user.(models.User)

    var quiz models.Quiz

    // Cari kuis berdasarkan ID
    if err := postgresql.DB.Preload("Questions").First(&quiz, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Quiz not found"})
        return
    }

    // Ambil jawaban pengguna
    var body struct {
        Answers []struct {
            QuestionID uint   `json:"question_id" binding:"required"`
            Answer     string `json:"answer" binding:"required"` // Jawaban pengguna (A, B, C, atau D)
        } `json:"answers" binding:"required,len=10"`
    }

    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
        return
    }

    // Periksa jawaban
    correctCount := 0
    for _, userAnswer := range body.Answers {
        for _, question := range quiz.Questions {
            if question.ID == userAnswer.QuestionID && question.CorrectAnswer == userAnswer.Answer {
                correctCount++
                break
            }
        }
    }

    // Hitung poin berdasarkan jumlah jawaban benar
    points := correctCount * 10
    currentUser.Points += points

    if err := postgresql.DB.Save(&currentUser).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user points"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message":       "Quiz completed successfully",
        "correct_count": correctCount,
        "points_awarded": points,
    })
}

