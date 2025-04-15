package main

import (
	"log"

	"github.com/888zxc/microfast/config"
	"github.com/888zxc/microfast/internal/server"
	"github.com/888zxc/microfast/internal/limiter"
	"github.com/888zxc/microfast/internal/logger"
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
