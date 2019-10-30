FROM golang:1.13.3-alpine AS builder

WORKDIR /

RUN apk add --no-cache git

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o service 

FROM alpine:latest

COPY --from=builder /service .

CMD ["/service"]