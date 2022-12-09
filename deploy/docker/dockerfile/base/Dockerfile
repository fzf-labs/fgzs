FROM alpine:latest
# RUN 设置代理镜像
RUN echo -e http://mirrors.ustc.edu.cn/alpine/v3.16/main/ > /etc/apk/repositories

# RUN 设置 Asia/Shanghai 时区
RUN apk --no-cache add tzdata  && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone \

WORKDIR /app
