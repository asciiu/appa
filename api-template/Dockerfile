FROM golang:1.13.3 AS builder
ARG SSH_PRIVATE_KEY

WORKDIR /

RUN mkdir -p ~/.ssh && umask 0077 && echo "${SSH_PRIVATE_KEY}" > ~/.ssh/id_rsa \
	&& git config --global url."git@github.com:".insteadOf https://github.com/ \
	&& ssh-keyscan github.com >> ~/.ssh/known_hosts

COPY api-template/ api-template
COPY lib/ lib 

RUN cd api-template && CGO_ENABLED=0 GOOS=linux go build -o api 

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /api-template/api .

EXPOSE 9000

CMD ["/api"]