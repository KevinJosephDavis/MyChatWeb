package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MySQLDSN      string
	RedisAddr     string
	RedisPassword string
	RedisDB       int
}

func LoadConfig() *Config {
	// 加载 .env 文件
	godotenv.Load()

	return &Config{
		MySQLDSN:      getEnv("MYSQL_DSN", "root:123456@tcp(localhost:3306)/chatroom?charset=utf8mb4&parseTime=True&loc=Local"),
		RedisAddr:     getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       0,
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
