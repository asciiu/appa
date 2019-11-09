FROM golang:1.13.3 AS builder
ARG SSH_PRIVATE_KEY

WORKDIR /

RUN mkdir -p ~/.ssh && umask 0077 && echo "${SSH_PRIVATE_KEY}" > ~/.ssh/id_rsa \
	&& git config --global url."git@github.com:".insteadOf https://github.com/ \
	&& ssh-keyscan github.com >> ~/.ssh/known_hosts

COPY youtube/ tube 
COPY lib/ lib 

RUN cd youtube && CGO_ENABLED=0 GOOS=linux go build -o gotube 

FROM alpine:latest

RUN apk --no-cache add ca-certificates ffmpeg

COPY --from=builder /youtube/gotube .

CMD ["/gotube"]