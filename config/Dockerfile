FROM alpine:latest as certs

# Install the CA certificates
RUN apk --update add ca-certificates


FROM scratch AS prod

ARG USER_UID=10001
USER ${USER_UID}

# 从certs阶段拷贝CA证书
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
# 从buil阶段拷贝二进制文件
COPY --from=build ../bin/otelcontribcol_linux_amd64 ./otelcontribcol

ENTRYPOINT["./otelcontribcol"]

CMD ["--config", "/etc/otel/relation-metrics-config.yaml"]