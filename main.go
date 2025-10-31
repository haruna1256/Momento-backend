package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haruna1256/Momento-backend/database"
	"github.com/haruna1256/Momento-backend/handlers"
)

func main() {
     // データベース接続
    database.ConnectDB()

    // ルーター作成
    router := gin.Default()

    // エンドポイント設定
    router.GET("/photos", handlers.GetPhotos)
    router.POST("/photos", handlers.CreatePhoto)
    fmt.Println("Server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
