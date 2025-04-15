package api

import (
	"net/http"
	"time"

	"github.com/888zxc/Go-server/internal/middleware"
)

// NewRouter 设置并返回HTTP路由器
func NewRouter() http.Handler {
	// 创建ServeMux实例
	mux := http.NewServeMux()

	// 注册路由
	mux.HandleFunc("/", handleHome)
	mux.HandleFunc("/api/health", handleHealth)
	mux.HandleFunc("/api/data", handleData)

	// 应用中间件
	var handler http.Handler = mux
	handler = middleware.Logger(handler)
	handler = middleware.Recovery(handler)
	handler = middleware.Cors(handler)
	handler = middleware.RateLimit(handler, 100) // 每秒最多处理100个请求

	return handler
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("欢迎访问高性能服务器!"))
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok","time":"` + time.Now().Format(time.RFC3339) + `"}`))
}

func handleData(w http.ResponseWriter, r *http.Request) {
	// 模拟数据处理
	if r.Method != http.MethodGet {
		http.Error(w, "仅支持GET请求", http.StatusMethodNotAllowed)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"data":[{"id":1,"name":"Item 1"},{"id":2,"name":"Item 2"}]}`))
}
