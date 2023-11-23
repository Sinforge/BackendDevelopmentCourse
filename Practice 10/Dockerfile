FROM golang:1.19.3-alpine AS builder

WORKDIR /appdir

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./

RUN go build -o main rest/server/main.go

FROM alpine

WORKDIR /appdir

COPY --from=builder /appdir .

EXPOSE 8080
CMD ["/appdir/main"]