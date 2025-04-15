package api

import (
	"encoding/json"
	"net/http"
	"runtime"
	"time"
)

// RegisterMetricsHandler 添加监控接口
func RegisterMetricsHandler(mux *http.ServeMux) {
	mux.HandleFunc("/metrics", handleMetrics)
}

func handleMetrics(w http.ResponseWriter, r *http.Request) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	metrics := map[string]interface{}{
		"time":         time.Now().Format(time.RFC3339),
		"goroutines":   runtime.NumGoroutine(),
		"memory": map[string]interface{}{
			"alloc":      m.Alloc,
			"total_alloc": m.TotalAlloc,
			"sys":        m.Sys,
			"heap_alloc": m.HeapAlloc,
			"heap_sys":   m.HeapSys,
		},
		"gc": map[string]interface{}{
			"next_gc": m.NextGC,
			"num_gc":  m.NumGC,
		},
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metrics)
}
