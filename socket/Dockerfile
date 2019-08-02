ARG GO_VERSION=1.12

# First stage: build the executable.
FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /

RUN apk add --no-cache git

COPY . .

RUN go mod download 

RUN CGO_ENABLED=0 GOOS=linux go build -o socket

FROM alpine:latest

COPY --from=builder /socket .

CMD ["./socket"]
