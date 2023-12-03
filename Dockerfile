# 使用golang的alpine版本作为基础镜像
FROM golang:alpine
# 设置工作目录
WORKDIR /app
# 将当前目录的内容复制到容器中的/app目录
COPY . .
# 设置Go模块支持，并设置Go代理
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io"
# 编译Go应用程序
RUN go build -o myapp main.go
# 暴露9988端口
EXPOSE 9988
# 设置容器启动时运行的命令
ENTRYPOINT ["./myapp"]