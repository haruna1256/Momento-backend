package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "github.com/haruna1256/Momento-backend/models"
    "github.com/haruna1256/Momento-backend/database"
)

// POST /photos
func CreatePhoto(c *gin.Context) {
    var photo models.Photo

    // リクエストBodyをJSONとして受け取る
    if err := c.ShouldBindJSON(&photo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // UUIDを自動生成
    photo.PhotoID = uuid.New()

    // DBに保存
    result := database.DB.Create(&photo)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, photo)
}

// GET /photos
func GetPhotos(c *gin.Context) {
    var photos []models.Photo
    result := database.DB.Find(&photos)

    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, photos)
}
