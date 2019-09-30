
# First stage: build the executable.
FROM golang:1.12-alpine AS builder

WORKDIR /

RUN apk update
RUN apk add --no-cache git gcc libc-dev libgit2-0.27-dev pkgconf 

COPY . . 

RUN CGO_ENABLED=1 GOOS=linux go build -o app 

FROM alpine:latest

RUN apk add --no-cache libgit2-0.27 

COPY --from=builder /app /service

# Run the compiled binary.
#ENTRYPOINT ["/service", "--selector=static", "--server_address=:8080", "--broker_address=:10001"]
CMD ["/service"]