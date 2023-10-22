# Слой с основным образом
FROM golang:latest as builder

WORKDIR /go-cookie-service
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/server/main.go

# Слой с образом для запуска
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /go-cookie-service .

CMD ["./app"]
