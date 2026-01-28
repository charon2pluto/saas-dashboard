# 1. Build Stage
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
# 编译 Go 程序
RUN go build -o main .

# 2. Run Stage (使用极小的 Alpine 镜像)
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
# 暴露 8080 端口
EXPOSE 8080
CMD ["./main"]