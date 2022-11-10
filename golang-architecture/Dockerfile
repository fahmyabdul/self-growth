FROM golang:1.19.2-bullseye AS builder

ENV GO111MODULE="auto"

RUN apt-get install git

WORKDIR /golang-architecture

COPY . .

RUN go mod vendor

FROM golang:1.19.2-bullseye

ENV GO111MODULE="auto"

WORKDIR /golang-architecture

COPY --from=builder /golang-architecture ./

USER root

EXPOSE 8500

CMD ["/bin/bash","/golang-architecture/scripts/run.sh"]