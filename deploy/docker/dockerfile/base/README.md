制作基础镜像

1.构建镜像

```shell
docker build -t go-zero-alpine:v1 .
```

2.打tag

```shell
docker tag go-zero-alpine:v1 120.77.97.206:8077/library/go-zero-alpine:v1

```

3.推送

```shell
docker push 120.77.97.206:8077/library/go-zero-alpine:v1
```

拉取

```shell
docker pull 120.77.97.206:8077/library/go-zero-alpine:v1
```
