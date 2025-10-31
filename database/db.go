package database

import (
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "github.com/haruna1256/Momento-backend/models"
)

var DB *gorm.DB

func ConnectDB() {
    // SQLiteを使ってDB接続
    database, err := gorm.Open(sqlite.Open("momento.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database:", err)
    }

    // モデルを自動マイグレーション（テーブルがなければ作成）
    database.AutoMigrate(&models.Photo{})

    DB = database
    log.Println("データベースが接続され移行された。")
}