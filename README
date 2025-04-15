# microfast

[![CI](https://github.com/888zxc/microfast/actions/workflows/ci.yml/badge.svg)](https://github.com/888zxc/microfast/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/888zxc/microfast)](https://goreportcard.com/report/github.com/888zxc/microfast)
[![license](https://img.shields.io/github/license/888zxc/microfast.svg)](./LICENSE)

## 项目简介

**microfast** 是一款为云原生和微服务架构打造的超高性能、高并发 Golang HTTP 服务器。  
它采用 fasthttp 内核、分层模块化设计，开箱即用地集成了日志、限流、安全、中间件、优雅关停、Prometheus 监控等现代服务必备能力，极易扩展业务路由与微服务协作。  
适用于高吞吐网关、API中间层、SaaS底座、AI服务接口等全部高并发场景。

---

## Features

- 🚀 **极致高性能**：fasthttp + 优化设计，百万级QPS不是梦
- 🛡️ **安全默认**：panic恢复、加固HTTP头、限流防护即插即用
- 📈 **可观测性**：Prometheus指标与健康探针，全链路易集成
- ♻️ **优雅关停**：支持平滑升级与弹性伸缩
- 🧩 **模块化易扩展**：中间件、API 路由、高可维护
- 🛠️ **工程友好**：CI/测试/lint示例全配齐

---

## 目录结构

```
microfast/
│
├── cmd/server/main.go           # 程序入口
├── config/config.go             # 配置模块
├── internal/
│   ├── middleware/              # 标准中间件
│   ├── handler/                 # 路由与业务入口
│   ├── server/                  # 服务启动
│   ├── limiter/                 # 限流实现
│   └── logger/                  # 日志封装
├── pkg/version/version.go       # 版本信息
├── scripts/bench.go             # 性能压测脚本
├── openapi.yaml                 # Swagger API文档
├── .github/workflows/ci.yml     # CI流水线
├── go.mod go.sum
└── README.md
```

---

## 快速开始

### 依赖环境

- Go 1.20+（或更高）
- [可选] make/curl/wrk 作为辅助

### 拉取&构建

```bash
git clone https://github.com/888zxc/microfast.git
cd microfast
go mod tidy
go build -o microfast-server ./cmd/server
```

### 启动服务器

```bash
# 默认8080端口
./microfast-server

# 可通过环境变量定制端口与限流
SERVER_PORT=9000 LIMIT_PER_SEC=20000 ./microfast-server
```

### 关键参数说明

- `SERVER_PORT`：监听端口（默认8080）
- `LIMIT_PER_SEC`：每秒限流（默认20000）

---

### 访问和测试API

```bash
curl http://localhost:8080/           # 首页欢迎
curl http://localhost:8080/healthz    # 健康检查
curl http://localhost:8080/metrics    # Prometheus指标
```

---

## Swagger & API 文档

详见 [openapi.yaml](./openapi.yaml)，可在 [Swagger Editor](https://editor.swagger.io/) 在线预览。

接口摘要：

| Path      | Method | 说明              |
|-----------|--------|-------------------|
| `/`       | GET    | 服务器欢迎页      |
| `/healthz`| GET    | 健康探针          |
| `/metrics`| GET    | 状态指标 (Prom)   |

如需扩展API，只需在 `internal/handler/handler.go` 编写新的路由和业务代码。

---

## 性能压测

### 推荐压测工具

**Go 内置压测脚本：**

```bash
go run scripts/bench.go
```

**第三方工具示例：**

```bash
# wrk
wrk -t8 -c1000 -d10s http://localhost:8080/

# hey
hey -n 100000 -c 500 http://localhost:8080/healthz
```

性能参考：i9物理机下10万QPS+，可按机器与实际需求调并发参数。

---

## 持续集成（GitHub Actions）

本仓库集成自动 CI，支持：

- 自动 Lint、Gofmt、单元测试、竞态检测、构建
- [ci.yml](.github/workflows/ci.yml) 可根据需要拓展容器构建与发布

---

## 常见问题

- 端口占用、权限问题：请换端口或用 sudo（谨慎！）
- 如何扩展Router：直接在 handler/handler.go 添加新分支即可
- Prometheus采集：在 targets 添加 `/metrics` 即可

---

## 贡献方式

欢迎提Star、Issue和PR！

1. Fork & 新建分支
2. 开发与测试
3. 提交 PR

---

## License

[MIT](./LICENSE)

---

## 致谢

- [fasthttp](https://github.com/valyala/fasthttp)
- [zap](https://github.com/uber-go/zap)
- [prometheus/client_golang](https://github.com/prometheus/client_golang)

---

如有疑问或需要企业支持，欢迎联系作者 issue！

---

> ⭐️ 如果本项目帮助到你，欢迎 Star 与分享！
