package controllers
import(
	"net/http"
	"trashure/internal/domain/entity"
	"trashure/internal/infra/postgresql"
    "fmt"
	"github.com/gin-gonic/gin"
)
// CreateTrashureRequest - User membuat permintaan pengumpulan sampah
func CreateTrashureRequest(c *gin.Context) {
    // Ambil tipe sampah dan berat dari form
    type TrashureRequestInput struct {
        Type   string  `form:"type" binding:"required"`
        Weight float64 `form:"weight" binding:"required"`
    }

    var input TrashureRequestInput
    if err := c.ShouldBind(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
        return
    }

    // Ambil file gambar dari request
    fileHeader, err := c.FormFile("image")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Image file is required"})
        return
    }

    // Unggah gambar ke Supabase dan dapatkan URL-nya
    imageURL, err := UploadImageToSupabase(fileHeader)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to upload image: %s", err.Error())})
        return
    }

    // Simpan URL gambar dan data lainnya ke database
    trashureRequest := entity.TrashureRequest{
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
    var trashureRequest entity.TrashureRequest

    // Cari Trashure Request berdasarkan ID
    if err := postgresql.DB.First(&trashureRequest, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Trashure request not found"})
        return
    }

    // Kembalikan data Trashure Request
    c.JSON(http.StatusOK, gin.H{"data": trashureRequest})
}

// ConfirmTrashureRequest - Admin mengkonfirmasi permintaan dan menetapkan harga
func ConfirmTrashureRequest(c *gin.Context) {
    id := c.Param("id")
    var body struct {
        Price float64 `json:"price" binding:"required"`
    }

    if err := c.Bind(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
        return
    }

    var request entity.TrashureRequest

    // Temukan permintaan
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
    var request entity.TrashureRequest

    // Cari permintaan berdasarkan ID
    if err := postgresql.DB.First(&request, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Trashure request not found"})
        return
    }

    // Hapus permintaan dari database
    if err := postgresql.DB.Delete(&request).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete trashure request"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Trashure request deleted successfully", "id": id})
}