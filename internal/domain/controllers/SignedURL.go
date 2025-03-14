package controllers

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
)

type SignedURLResponse struct {
    SignedURL string `json:"signedURL"`
}

// CreateSignedURL - Membuat Signed URL untuk file di Supabase
func CreateSignedURL(filePath string, expiresIn int) (string, error) {
    // Supabase configuration
    supabaseURL := os.Getenv("SUPABASE_URL")        
    supabaseAPIKey := os.Getenv("SUPABASE_API_KEY") 
    bucket := os.Getenv("SUPABASE_BUCKET")          

    // Endpoint untuk membuat Signed URL
    url := fmt.Sprintf("%s/storage/v1/object/sign/%s/%s", supabaseURL, bucket, filePath)

    // Payload untuk request Signed URL
    payload := map[string]interface{}{
        "expiresIn": expiresIn, 
    }

    // Serialisasi payload ke JSON
    jsonData, err := json.Marshal(payload)
    if err != nil {
        return "", fmt.Errorf("failed to marshal payload: %w", err)
    }

    // Membuat permintaan HTTP POST
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
    if err != nil {
        return "", fmt.Errorf("failed to create request: %w", err)
    }

    // Menambahkan header
    req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", supabaseAPIKey))
    req.Header.Set("Content-Type", "application/json")

    // Kirim permintaan
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", fmt.Errorf("failed to send request: %w", err)
    }
    defer resp.Body.Close()

    // Periksa status respons
    if resp.StatusCode != http.StatusOK {
        body, _ := ioutil.ReadAll(resp.Body)
        return "", fmt.Errorf("failed to create signed URL: %s", string(body))
    }

    // Parsing respons JSON
    var signedURLResponse SignedURLResponse
    if err := json.NewDecoder(resp.Body).Decode(&signedURLResponse); err != nil {
        return "", fmt.Errorf("failed to decode response: %w", err)
    }

    return signedURLResponse.SignedURL, nil
}
func GetSignedURL(c *gin.Context) {
    // Ambil path file dari query (contoh: "images/example.jpg")
    filePath := c.Query("file")
    if filePath == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "File path is required"})
        return
    }

    // Waktu kadaluarsa Signed URL dalam detik (contoh: 3600 detik = 1 jam)
    expiresIn := 3600

    // Buat Signed URL
    signedURL, err := CreateSignedURL(filePath, expiresIn)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Kembalikan Signed URL ke pengguna
    c.JSON(http.StatusOK, gin.H{
        "signed_url": signedURL,
    })
}
