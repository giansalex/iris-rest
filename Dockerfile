FROM golang:1.14-alpine AS builder

WORKDIR /root
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -a -o /root/app

FROM scratch
LABEL owner="Giancarlos Salas"
LABEL maintainer="giansalex@gmail.com"
EXPOSE 8080

WORKDIR /root
COPY --from=builder /root/app .

ENTRYPOINT ["./app"]
