package main

import (
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

    // サーバー起動
    router.Run(":8080") // ここでサーバーを起動
}
