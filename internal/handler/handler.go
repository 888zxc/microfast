package handler

import (
	"fmt"
	"time"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	reqTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path", "method"},
	)
	reqDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path"},
	)
)

func init() {
	prometheus.MustRegister(reqTotal, reqDuration)
}

// 创建 metrics handler 的 fasthttp 适配器
var metricsHandler = fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler())

func MainHandler(ctx *fasthttp.RequestCtx) {
    start := time.Now()
    p := string(ctx.Path())
    m := string(ctx.Method())

    switch p {
    case "/healthz":
        ctx.SetStatusCode(fasthttp.StatusOK)
        ctx.SetBodyString("OK")
    case "/":
        // 这里你的处理逻辑
    case "/metrics":
        // 直接调用转换后的 handler
        metricsHandler(ctx)
    default:
        ctx.SetStatusCode(fasthttp.StatusNotFound)
        ctx.SetBodyString("not found")
    }

    reqTotal.WithLabelValues(p, m).Inc()
    reqDuration.WithLabelValues(p).Observe(time.Since(start).Seconds())
	fmt.Printf("Path: %s, Status: %d\n", ctx.Path(), ctx.Response.StatusCode())
}
