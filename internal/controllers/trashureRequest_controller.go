package controllers
import(
	"net/http"
	"trashure/internal/models"
	"trashure/internal/postgresql"
    "fmt"
	"github.com/gin-gonic/gin"
)

func CreateTrashureRequest(c *gin.Context) {
    type TrashureRequestInput struct {
        Type   string  `form:"type" binding:"required"`
        Weight float64 `form:"weight" binding:"required"`
    }

    var input TrashureRequestInput
    if err := c.ShouldBind(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
        return
    }

    fileHeader, err := c.FormFile("image")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Image file is required"})
        return
    }

    imageURL, err := UploadImageToSupabase(fileHeader)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to upload image: %s", err.Error())})
        return
    }

    trashureRequest := models.TrashureRequest{
        Type:     input.Type,
        Weight:   input.Weight,
        ImageURL: imageURL,
        Status:   "pending",
    }

    if err := postgresql.DB.Create(&trashureRequest).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save trashure request"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Trashure request created successfully",
        "data":    trashureRequest,
    })
}

func GetTrashureRequest(c *gin.Context) {
    id := c.Param("id")
    var trashureRequest models.TrashureRequest

    if err := postgresql.DB.First(&trashureRequest, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Trashure request not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": trashureRequest})
}

func ConfirmTrashureRequest(c *gin.Context) {
    id := c.Param("id")
    var body struct {
        Price float64 `json:"price" binding:"required"`
    }

    if err := c.Bind(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
        return
    }

    var request models.TrashureRequest

    
    if err := postgresql.DB.First(&request, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Trashure request not found"})
        return
    }

    request.Price = body.Price
    request.Status = "confirmed"

    if err := postgresql.DB.Save(&request).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to confirm trashure request"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Trashure request confirmed", "data": request})
}
func DeleteTrashureRequest(c *gin.Context) {
    id := c.Param("id")
    var request models.TrashureRequest

    if err := postgresql.DB.First(&request, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Trashure request not found"})
        return
    }

    if err := postgresql.DB.Delete(&request).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete trashure request"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Trashure request deleted successfully", "id": id})
}