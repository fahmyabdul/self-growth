FROM golang:1.19.2-bullseye AS builder

RUN apt-get install git

WORKDIR /data-app

COPY . .

RUN go mod vendor

FROM golang:1.19.2-bullseye

ENV GOVERSION="go1.19.2" \
    GO111MODULE="auto" \
    GOCACHE=/tmp \
    XDG_CACHE_HOME=/tmp/.parameter

WORKDIR /data-app

COPY --from=builder /data-app ./

USER root

RUN apt-get update && \
    apt-get install -y nano && \
    chmod -R 777 /data-app

EXPOSE 8600

CMD ["/bin/bash","/data-app/scripts/run.sh"]