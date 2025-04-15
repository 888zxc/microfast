package handler

import (
	"time"
	"github.com/valyala/fasthttp"
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

// 路由主入口
func MainHandler(ctx *fasthttp.RequestCtx) {
	start := time.Now()
	p := string(ctx.Path())
	m := string(ctx.Method())

	switch p {
	case "/healthz":
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBodyString("OK")
	case "/":
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBodyString("Welcome to microfast!")
	case "/metrics":
		promhttp.InstrumentMetricHandler(
			prometheus.DefaultRegisterer, promhttp.Handler(),
		).ServeHTTP(ctx.Response.BodyWriter(), ctx.RequestCtx.Request)
	default:
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		ctx.SetBodyString("not found")
	}

	reqTotal.WithLabelValues(p, m).Inc()
	reqDuration.WithLabelValues(p).Observe(time.Since(start).Seconds())
}
