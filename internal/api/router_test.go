package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkHandleData(b *testing.B) {
	// 创建测试服务器
	handler := http.HandlerFunc(handleData)
	server := httptest.NewServer(handler)
	defer server.Close()
	
	// 执行基准测试
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := http.Get(server.URL)
		if err != nil {
			b.Fatal(err)
		}
		resp.Body.Close()
	}
}
