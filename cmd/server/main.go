package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/888zxc/Go-server/internal/api"
	"github.com/888zxc/Go-server/internal/config"
)

func main() {
	// 加载配置
	cfg := config.Load()

	// 创建路由器
	router := api.NewRouter()

	// 配置HTTP服务器
	server := &http.Server{
		Addr:         cfg.ServerAddress,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// 在单独的goroutine中启动服务器
	go func() {
		log.Printf("服务器启动于 %s\n", cfg.ServerAddress)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("服务器启动失败: %v\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("正在关闭服务器...")

	// 设置关闭超时
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("服务器强制关闭: %v\n", err)
	}

	log.Println("服务器优雅退出")
}
