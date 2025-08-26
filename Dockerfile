# 使用官方的Golang镜像作为基础镜像
FROM golang:1.19

# 设置工作目录
WORKDIR /app

# 将本地文件复制到容器中
COPY . .

# 使用Go Modules下载依赖并编译项目
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go mod download \
    && go mod tidy \
    && go build -o server .

# 暴露端口
EXPOSE 8080

# 启动应用程序
CMD ["./server"]
