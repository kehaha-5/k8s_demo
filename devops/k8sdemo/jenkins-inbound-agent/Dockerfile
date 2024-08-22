FROM jenkins/inbound-agent:alpine3.20-jdk17

# 切换到 root 用户
USER root

RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && apk update  &&\
apk add nodejs npm && npm config set registry https://registry.npmmirror.com/ && \
npm install @vue/cli@5.0.8

COPY go1.22.6.linux-amd64.tar.gz /app/go1.22.6.linux-amd64.tar.gz
COPY kubectl /usr/local/bin/

RUN rm -rf /usr/local/go && \
    tar -C /usr/local -xzf /app/go1.22.6.linux-amd64.tar.gz && \
    ln -s /usr/local/go/bin/go /usr/local/bin/go && rm -rf /app/go1.22.6.linux-amd64.tar.gz

USER jenkins
