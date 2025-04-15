package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/reuseport"
	"go.uber.org/zap"


	"github.com/888zxc/microfast/internal/handler"
	"github.com/888zxc/microfast/internal/middleware"
	"github.com/888zxc/microfast/internal/limiter"
	"github.com/888zxc/microfast/internal/logger"
)

type Server struct {
	server  *fasthttp.Server
	port    string
}

func NewServer(port string, limiter *limiter.Limiter) *Server {
	handlerChain := middleware.Chain(
		limiter,
		middleware.Recovery(),
		middleware.Logging(),
		middleware.SecureHeaders(),
	)(handler.MainHandler)

	s := &fasthttp.Server{
		Handler:            handlerChain,
		Name:               "microfast",
		Logger:             log.New(os.Stderr, "[microfast] ", log.LstdFlags),
		ReadTimeout:        30 * time.Second,
		WriteTimeout:       30 * time.Second,
		MaxConnsPerIP:      10000,
		MaxRequestsPerConn: 1000,
	}
	return &Server{
		server:  s,
		port:    port,
	}
}

func (s *Server) Start() error {
	ln, err := reuseport.Listen("tcp4", ":"+s.port)
	if err != nil {
		return err
	}

	go func() {
		logger.L().Info("Server listening", zap.String("port", s.port))
		if err := s.server.Serve(ln); err != nil {
			logger.L().Error("Server Serve error", zap.Error(err))
		}
	}()

	// 优雅关停
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	logger.L().Info("Shutting down server gracefully")
	return s.server.Shutdown()
}
