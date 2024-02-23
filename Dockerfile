# 使用多阶段构建，首先使用golang官方镜像作为构建环境
FROM golang:1.18 as builder

# 设置工作目录
WORKDIR /build

# 将代码复制到/build目录下
COPY . .

# 编译项目，生成二进制文件
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o main .

# 使用scratch作为基础镜像，创建最终的运行环境
FROM scratch

# 创建/data目录，并将二进制文件复制到/data目录下
WORKDIR /data
COPY --from=builder /build/main .
COPY --from=builder /build/static ./static

# 开放8228端口
EXPOSE 8228

# 运行二进制文件
CMD ["./main"]