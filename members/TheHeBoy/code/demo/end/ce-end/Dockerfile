# 构建阶段
FROM golang:alpine AS builder
# 构建可执行文件
ENV CGO_ENABLED=0
WORKDIR /build
ADD . .
RUN go mod tidy
RUN go build -o main

# 运行阶段
FROM alpine
# 安装 tzdata 包以设置时区
RUN apk add --no-cache tzdata
# 设置时区为上海
ENV TZ=Asia/Shanghai
# 复制时区信息到系统时区文件
RUN cp /usr/share/zoneinfo/${TZ} /etc/localtime && \
    echo "${TZ}" > /etc/timezone
# 设置工作目录
WORKDIR /app
# 复制构建阶段的文件
COPY --from=builder /build/application.yaml .
COPY --from=builder /build/application-local.yaml .
COPY --from=builder /build/application-test.yaml .
COPY --from=builder /build/storage/template ./storage/template
COPY --from=builder /build/main .
# 设置编码
ENV LANG=C.UTF-8
# 暴露端口和卷
EXPOSE 3000
VOLUME /app/storage/logs
# 设置容器启动命令
CMD ["./main"]
