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
    router.GET("/photos", handlers.GetPhotos)      // 写真一覧取得
    router.POST("/photos", handlers.CreatePhoto)   // 写真情報を登録（URLなど）
    router.POST("/upload", handlers.UploadPhoto)   // 実際の画像ファイルをアップロード

    // サーバー起動
    router.Run(":4000") // ここでサーバーを起動
}
