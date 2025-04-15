package middleware

import (
	"github.com/valyala/fasthttp"
	"github.com/888zxc/microfast/internal/logger"
	"github.com/888zxc/microfast/internal/limiter"
	"go.uber.org/zap"
)

type Middleware func(next fasthttp.RequestHandler) fasthttp.RequestHandler

// Panic恢复 + 日志 + 安全头 + 限流
func Chain(limiter *limiter.Limiter, handlers ...Middleware) Middleware {
	return func(final fasthttp.RequestHandler) fasthttp.RequestHandler {
		// 包裹中间件
		for i := len(handlers) - 1; i >= 0; i-- {
			final = handlers[i](final)
		}
		return func(ctx *fasthttp.RequestCtx) {
			// 限流
			if !limiter.Allow() {
				ctx.SetStatusCode(fasthttp.StatusTooManyRequests)
				ctx.SetBodyString("Too Many Requests")
				return
			}
			final(ctx)
		}
	}
}

func Logging() Middleware {
	return func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			logger.L().Info("Request",
				zap.String("method", string(ctx.Method())),
				zap.String("path", string(ctx.Path())),
				zap.String("remote_addr", ctx.RemoteAddr().String()),
			)
			next(ctx)
		}
	}
}

func Recovery() Middleware {
	return func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			defer func() {
				if r := recover(); r != nil {
					logger.L().Error("panic", zap.Any("error", r))
					ctx.SetStatusCode(fasthttp.StatusInternalServerError)
					ctx.SetBodyString("Internal Server Error")
				}
			}()
			next(ctx)
		}
	}
}

func SecureHeaders() Middleware {
	return func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			ctx.Response.Header.Set("X-Content-Type-Options", "nosniff")
			ctx.Response.Header.Set("X-Frame-Options", "DENY")
			ctx.Response.Header.Set("X-XSS-Protection", "1; mode=block")
			next(ctx)
		}
	}
}
