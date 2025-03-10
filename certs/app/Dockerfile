FROM swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/library/golang:1.24 AS builder
COPY ./go.mod ./go.mod 
COPY ./main.go ./main.go

RUN go build -o /app ./main.go


FROM swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/library/debian:11-slim

COPY --from=builder --chown=user:group /app /app

USER 1001

CMD ["/app"]