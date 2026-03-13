package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Load 從 .env 載入環境變數
func Load() error {
	return godotenv.Load()
}

// Get 取得環境變數，若無則回傳預設值
func Get(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}

// Port 回傳 API 監聽的 port
func Port() string {
	return Get("PORT", "8080")
}

// DatabaseURL 回傳 PostgreSQL 連線字串
// 格式：postgres://user:password@host:port/dbname?sslmode=disable
func DatabaseURL() string {
	return Get("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/practice?sslmode=disable")
}
