ARG GO_VERSION=1.13.3

# First stage: build the executable.
FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /

RUN apk add --no-cache git

COPY . .

RUN go mod download 

RUN CGO_ENABLED=0 GOOS=linux go build -o api 

FROM alpine:latest

COPY --from=builder /api .

EXPOSE 9000

CMD ["/api"]