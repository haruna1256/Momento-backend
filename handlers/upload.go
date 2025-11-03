package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func UploadPhoto(c *gin.Context) {
	// 環境変数読み込み
	godotenv.Load()
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	bucket := os.Getenv("SUPABASE_BUCKET")

	// フォームデータからファイル取得
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ファイルを取得できません"})
		return
	}

	// ファイルを開く
	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ファイルを開けません"})
		return
	}
	defer file.Close()

	// ファイルデータをバイナリに変換
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ファイルの読み込みに失敗"})
		return
	}

	fileName := filepath.Base(fileHeader.Filename)

	// Supabase Storage API へのリクエスト準備
	url := fmt.Sprintf("%s/storage/v1/object/%s/%s", supabaseURL, bucket, fileName)

	req, err := http.NewRequest("POST", url, bytes.NewReader(fileBytes))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "リクエスト作成に失敗"})
		return
	}

	req.Header.Set("Authorization", "Bearer "+supabaseKey)
	req.Header.Set("Content-Type", "application/octet-stream")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Supabaseへのアップロードに失敗"})
		return
	}
	defer resp.Body.Close()

	// エラーハンドリング
	if resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		c.JSON(resp.StatusCode, gin.H{"error": string(body)})
		return
	}

	// 公開URLを返す
	publicURL := fmt.Sprintf("%s/storage/v1/object/public/%s/%s", supabaseURL, bucket, fileName)
	c.JSON(http.StatusOK, gin.H{
		"message":   "アップロード成功",
		"file_name": fileName,
		"url":       publicURL,
	})
}
