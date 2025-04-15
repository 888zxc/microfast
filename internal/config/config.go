package config

import (
	"os"
	"strconv"
)

// Config 存储所有服务器配置
type Config struct {
	ServerAddress   string
	ReadTimeoutSec  int
	WriteTimeoutSec int
	IdleTimeoutSec  int
	MaxConns        int
}

// Load 从环境变量加载配置
func Load() *Config {
	return &Config{
		ServerAddress:   getEnv("SERVER_ADDR", ":8080"),
		ReadTimeoutSec:  getEnvAsInt("READ_TIMEOUT", 15),
		WriteTimeoutSec: getEnvAsInt("WRITE_TIMEOUT", 15),
		IdleTimeoutSec:  getEnvAsInt("IDLE_TIMEOUT", 60),
		MaxConns:        getEnvAsInt("MAX_CONNS", 1000),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
