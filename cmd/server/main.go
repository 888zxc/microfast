package main

import (
	"log"

	"github.com/microfast/config"
	"github.com/microfast/internal/server"
	"github.com/microfast/internal/limiter"
	"github.com/microfast/internal/logger"
)

func main() {
	conf := config.LoadConfig()
	logger.InitLogger()
	lim := limiter.NewLimiter(conf.LimitPerSec)
	srv := server.NewServer(conf.Port, lim)
	if err := srv.Start(); err != nil {
		log.Fatalf("Server exited with error: %v", err)
	}
}
