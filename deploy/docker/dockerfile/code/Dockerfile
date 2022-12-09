#自制镜像
FROM 120.77.97.206:8077/library/go-zero-alpine:v1

WORKDIR /app

COPY app/service/identity/rpc/etc /app/etc
COPY app/service/identity/rpc/identity-rpc /app/identity

CMD ["./identity-rpc", "-f", "etc/identity.yaml"]