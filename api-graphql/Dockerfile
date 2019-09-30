FROM golang:1.12-alpine AS builder

WORKDIR /

RUN apk add --no-cache git

COPY . . 

RUN CGO_ENABLED=0 GOOS=linux go build -o app 

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app /api

EXPOSE 8080

ENTRYPOINT ["/api"]