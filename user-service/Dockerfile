# We use the official golang image, which contains all the 
# correct build tools and libraries. Notice `as builder`,
# this gives this container a name that we can reference later on. 
FROM golang:1.11.3 as builder

# Set our workdir to our current service in the gopath
WORKDIR /go/src/github.com/asciiu/appa/user-service

ADD id_rsa /root/.ssh/id_rsa
RUN chmod 700 /root/.ssh/id_rsa
RUN echo "Host github.com\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config
RUN git config --global url.ssh://git@github.com/.insteadOf https://github.com/

ADD user-service /go/src/github.com/asciiu/appa/user-service
CMD cd /go/src/github.com/asciiu/appa/user-service && go get github.com/asciiu/appa && go build -o /appa

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure -vendor-only
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o user-service

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/asciiu/appa/user-service/user-service .

CMD ["./user-service"]
