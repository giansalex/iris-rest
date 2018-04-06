FROM golang:1.10-alpine AS builder

ENV DEP_VERSION 0.4.1

RUN apk update && apk add git && apk add curl
RUN curl -L -s https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 -o $GOPATH/bin/dep && \
    chmod +x $GOPATH/bin/dep

WORKDIR $GOPATH/src/github.com/giansalex/iris-rest
COPY . .
RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -a -installsuffix cgo -o /root/app

FROM alpine:latest
EXPOSE 8080

WORKDIR /root/
COPY --from=builder /root/app .

ENTRYPOINT ["./app"]