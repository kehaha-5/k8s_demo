FROM harhor.k8s.demo/k8s_demo/alpine@sha256:eddacbc7e24bf8799a4ed3cdcfa50d4b88a323695ad80f317b6629883b2c2a78
COPY ./k8s_demo /app/k8s_demo
COPY ./config-prod.yaml /app/config-prod.yaml
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk update && apk add tzdata