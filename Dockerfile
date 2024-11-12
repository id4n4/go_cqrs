ARG GO_VERSION=1.23.3

FROM golang:${GO_VERSION}-alpine AS builder

RUN go env -w GOPROXY=direct
RUN apk add --no-cache git
RUN apk add --no-cache ca-certificates && update-ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/planets-service ./cmd/cqrs-go/commands/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/query-service ./cmd/cqrs-go/querys/main.go

FROM alpine:3.18
WORKDIR /usr/bin

COPY --from=builder /go/bin/planets-service .
COPY --from=builder /go/bin/query-service .
