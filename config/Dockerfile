# BUILD 阶段

FROM golang:1.10 AS build

# 设置我们应用程序的工作目录
WORKDIR /go/src/github.com/opentelemetry-collector-hyperos
# 添加所有需要编译的应用代码
ADD ../. .

# 编译一个静态的go应用（在二进制构建中包含C语言依赖库）
RUN GO111MODULE=on CGO_ENABLED=0 go build  -o ./bin/otelcontribcol_linux_amd64   ./cmd/otelcontribcol

# 设置我们应用程序的启动命令
# ENTRYPOINT ["./bin/otelcontribcol_linux_amd64"]
# CMD ["--config", "/etc/otel/config.yaml"]



FROM alpine:latest as prep
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk --update add ca-certificates

RUN mkdir -p /tmp

FROM scratch

ARG USER_UID=10001
USER ${USER_UID}


# 从certs阶段拷贝CA证书
COPY --from=prep /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
# 从buil阶段拷贝二进制文件
COPY --from=build /go/src/github.com/opentelemetry-collector-hyperos/bin/otelcontribcol_linux_amd64 /otelcontribcol
ENTRYPOINT ["/otelcontribcol"]
CMD ["--config", "/etc/otel/config.yaml"]