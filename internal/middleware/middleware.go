package middleware

import (
	"log"
	"net"
	"net/http"
	"runtime/debug"
	"time"

	"golang.org/x/time/rate"
)

// Logger 中间件记录请求日志
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf(
			"%s - %s %s %s - %v",
			r.RemoteAddr,
			r.Method,
			r.URL.Path,
			r.Proto,
			time.Since(start),
		)
	})
}

// Recovery 中间件处理panic
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %v\n%s", err, debug.Stack())
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// Cors 中间件添加CORS头
func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}

// 并发控制
var limiter = rate.NewLimiter(100, 200) // 100 req/s, 突发200个请求

// RateLimit 实现限流中间件
func RateLimit(next http.Handler, rps int) http.Handler {
	limiter = rate.NewLimiter(rate.Limit(rps), rps*2)
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "请求过多，请稍后再试", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
