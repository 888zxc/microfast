# 选择极小体积的Go构建镜像
FROM golang:1.21-alpine AS builder

WORKDIR /app

# 拷贝源码并拉取依赖
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# 编译静态二进制文件（CGO=0），减小镜像体积
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o microfast-server ./cmd/server

# 裁剪镜像，仅包含可执行文件（二进制裸运行）
FROM scratch

WORKDIR /root/
COPY --from=builder /app/microfast-server .
COPY --from=builder /app/openapi.yaml ./   # 可选，挂载swagger文档

EXPOSE 8080

# 支持自定义端口，运行镜像可自行覆盖
ENV SERVER_PORT=8080
ENV LIMIT_PER_SEC=20000

ENTRYPOINT ["./microfast-server"]
