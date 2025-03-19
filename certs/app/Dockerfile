FROM swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/library/golang:1.24 AS builder
WORKDIR /app
COPY ./go.mod ./go.mod 
COPY ./main.go ./main.go
ENV CGO_ENABLED=1 
RUN go build -o https-client ./main.go
RUN chmod +x /app/https-client

FROM docker.empty233.top/library/debian:stable AS certs
RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list.d/debian.sources && \
    apt-get update && \
    apt-get install -y ca-certificates 

FROM docker.empty233.top/library/debian:stable
USER 1001
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder --chown=user:group /app/https-client /app/https-client
WORKDIR /app
CMD ["/app/https-client"]