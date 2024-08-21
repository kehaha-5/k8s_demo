# FROM alpine:3.20.2
FROM harhor.k8s.demo/k8s_demo/alpine-go:3.20.2
COPY ./k8s_demo /app/k8s_demo
COPY ./config-prod.yaml /app/config-prod.yaml
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk update && apk --no-cache add tzdata