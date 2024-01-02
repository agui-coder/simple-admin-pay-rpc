FROM alpine:3.19

# Define the project name | 定义项目名称
ARG PROJECT=pay
# Define the config file name | 定义配置文件名
ARG CONFIG_FILE=pay.yaml
# Define the author | 定义作者
ARG AUTHOR="894784649@qq.com"

LABEL org.opencontainers.image.authors=${AUTHOR}

WORKDIR /app
ENV PROJECT=${PROJECT}
ENV CONFIG_FILE=${CONFIG_FILE}

ENV TZ=Asia/Shanghai
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update --no-cache && apk add --no-cache tzdata

COPY ./${PROJECT}_rpc ./
COPY ./etc/${CONFIG_FILE} ./etc/
COPY ./cert/ ./cert/

EXPOSE 9007

ENTRYPOINT ./${PROJECT}_rpc -f etc/${CONFIG_FILE}