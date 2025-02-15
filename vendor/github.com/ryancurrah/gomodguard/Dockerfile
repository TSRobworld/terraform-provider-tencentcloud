# ---- Build container
FROM golang:1.24.0-alpine3.21 AS builder
WORKDIR /gomodguard
COPY . .
RUN apk add --no-cache git
RUN go build -o gomodguard cmd/gomodguard/main.go

# ---- App container
FROM golang:1.24.0-alpine3.21
WORKDIR /
RUN apk --no-cache add ca-certificates
COPY --from=builder gomodguard/gomodguard /
ENTRYPOINT ./gomodguard
