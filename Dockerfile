FROM alpine

# 创建/data目录，并将二进制文件复制到/data目录下
WORKDIR /data
COPY ./main .
COPY ./static ./static

# 开放8228端口
EXPOSE 8228

# 运行二进制文件
CMD ["./main"]
