FROM golang:1.11.3 as builder

WORKDIR /go/src/github.com/asciiu/appa/api

ADD id_rsa /root/.ssh/id_rsa
RUN chmod 700 /root/.ssh/id_rsa
RUN echo "Host github.com\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config
RUN git config --global url.ssh://git@github.com/.insteadOf https://github.com/

ADD api /go/src/github.com/asciiu/appa/api
CMD cd /go/src/github.com/asciiu/appa/api && go get github.com/asciiu/appa && go build -o /appa

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure -vendor-only
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/asciiu/appa/api/api .

EXPOSE 9000

CMD ["./api"]