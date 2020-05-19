FROM golang:1.14-alpine AS builder

WORKDIR /

RUN apk add --no-cache git

COPY api-graphql/ api-graphql 
COPY lib/ lib 

RUN cd api-graphql && CGO_ENABLED=0 GOOS=linux go build -o app 

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /api-graphql/app /api

EXPOSE 8080

ENTRYPOINT ["/api"]