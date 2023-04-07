FROM golang:1.18-bullseye

RUN go install github.com/beego/bee/v2@latest

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

ENV APP_HOME /go/src/mathapp

RUN mkdir -p "${APP_HOME}"

RUN go mod vendor

WORKDIR "${APP_HOME}"
EXPOSE 8080
CMD [ "bee", "run" ]
