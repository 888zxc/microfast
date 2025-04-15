package config

import (
	"os"
	"fmt"
)

type Config struct {
	Port       string
	LimitPerSec int
}

func LoadConfig() *Config {
	return &Config{
		Port:       getEnv("SERVER_PORT", "8080"),
		LimitPerSec: getEnvInt("LIMIT_PER_SEC", 20000),
	}
}

func getEnv(key, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}

func getEnvInt(key string, fallback int) int {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	var i int
	_, err := fmt.Sscanf(v, "%d", &i)
	if err != nil {
		return fallback
	}
	return i
}
