# First stage: build the executable.
FROM golang:1.12-alpine AS builder

WORKDIR /

RUN apk add --no-cache git

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o socket

FROM alpine:latest

COPY --from=builder /socket .

CMD ["./socket"]
