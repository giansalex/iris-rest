FROM golang:1.10 AS Builder

EXPOSE 8080
ENV DEP_VERSION 0.4.1

RUN curl -L -s https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 -o $GOPATH/bin/dep && \
    chmod +x $GOPATH/bin/dep

WORKDIR $GOPATH/src/github.com/giansalex/iris-rest
COPY . .
RUN dep ensure
RUN go build -o ./app .

ENTRYPOINT ["./app"]