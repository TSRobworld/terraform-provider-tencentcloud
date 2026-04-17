# ---- Build container
FROM golang:1.26.2-alpine3.23 AS builder
WORKDIR /gomodguard
COPY . .
RUN apk add --no-cache git
RUN go build -o gomodguard cmd/gomodguard/main.go

# ---- App container
FROM golang:1.26.2-alpine3.23
WORKDIR /
RUN apk --no-cache add ca-certificates
COPY --from=builder gomodguard/gomodguard /
ENTRYPOINT ./gomodguard
