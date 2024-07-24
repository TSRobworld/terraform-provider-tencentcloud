# ---- Build container
FROM golang:1.23rc2-alpine3.20 AS builder
WORKDIR /gomodguard
COPY . .
RUN apk add --no-cache git
RUN go build -o gomodguard cmd/gomodguard/main.go

# ---- App container
FROM golang:1.23rc2-alpine3.20
WORKDIR /
RUN apk --no-cache add ca-certificates
COPY --from=builder gomodguard/gomodguard /
ENTRYPOINT ./gomodguard
