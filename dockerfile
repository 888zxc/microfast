FROM golang:1.21-alpine AS builder

WORKDIR /app

# 复制Go模块文件
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 编译
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server ./cmd/server

# 最终镜像
FROM alpine:latest

RUN apk --no-cache add ca-certificates && \
    update-ca-certificates

WORKDIR /app
COPY --from=builder /app/server .

# 非root用户运行
RUN adduser -D appuser
USER appuser

EXPOSE 8080
CMD ["./server"]
