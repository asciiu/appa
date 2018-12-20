FROM golang:1.10.2 as builder

WORKDIR /go/src/github.com/asciiu/appa/socket
COPY . .
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o socket
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/asciiu/appa/socket/socket .

CMD ["./socket"]
