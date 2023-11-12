FROM golang:1.21 as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o main main.go

FROM debian:latest

WORKDIR /root

COPY --from=builder /app/main .

EXPOSE 3001

CMD ["./main"]