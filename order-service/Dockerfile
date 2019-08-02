ARG GO_VERSION=1.12

# First stage: build the executable.
FROM golang:${GO_VERSION}-alpine AS builder

# Set our workdir to our current service in the gopath
WORKDIR /

#ADD id_rsa /root/.ssh/id_rsa
#RUN chmod 700 /root/.ssh/id_rsa
#RUN echo "Host github.com\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config
#RUN git config --global url.ssh://git@github.com/.insteadOf https://github.com/

RUN apk add --no-cache git

COPY . .

RUN go mod download 

RUN CGO_ENABLED=0 GOOS=linux go build -o service 

FROM alpine:latest

COPY --from=builder /service .

CMD ["/service"]