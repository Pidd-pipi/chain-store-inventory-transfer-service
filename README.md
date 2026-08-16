# StockLink 门店库存调配后端服务

连锁零售企业门店库存运营平台的后端 API 服务，提供多门店 SKU 统一管理、库存同步预警、调拨审批、出入库记录与盘点、滞销分析与补货建议等能力。

## 技术栈

- Go 1.22
- Gin + GORM
- PostgreSQL 15
- JWT（golang-jwt/v5）+ RBAC

## 标准命令

```bash
go build ./...        # 编译
go test ./...         # 运行测试
go run ./cmd/server   # 启动 HTTP 服务
```

## 环境变量

| 变量 | 说明 | 默认值 |
| --- | --- | --- |
| DB_HOST | PostgreSQL 主机 | localhost |
| DB_PORT | PostgreSQL 端口 | 5432 |
| DB_NAME | 数据库名 | ldstoreinventory_db |
| DB_USER | 数据库用户 | ldstoreinventory_user |
| DB_PASSWORD | 数据库密码 | ldstoreinventory_pwd |
| JWT_SECRET | JWT 签名密钥 | change_me_to_a_long_random_string |
| TOKEN_TTL_HOURS | Token 有效期（小时） | 72 |
| APP_CORS_ORIGINS | 允许跨域来源（逗号分隔） | http://localhost:28504 |

## 目录结构

```
backend/
├── cmd/server/main.go
└── internal/
    ├── config/       # 配置解析
    ├── model/        # 实体定义（user/store/sku/store_inventory/transfer_order/stock_record/stocktake）
    ├── repository/   # 数据访问层
    ├── service/      # 业务逻辑层
    ├── handler/      # HTTP 接口层
    ├── router/       # 路由注册
    ├── middleware/   # auth/rbac/rate_limiter/error_handler/request_logger
    ├── dto/          # 请求/响应结构体
    ├── constants/    # 枚举、错误码、日志模板、文案
    └── util/         # jwt/logger/formatters/app_error/replenish_calculator
```

## 测试说明

业务逻辑层使用内存 mock 仓储完成单元测试，可在无外部数据库的环境直接运行 `go test ./...`。启动完整 HTTP 服务需要先准备 PostgreSQL 并设置上述环境变量。
