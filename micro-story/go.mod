module github.com/asciiu/appa/micro-story

go 1.13.3

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.6.0

replace github.com/nats-io/nats.go v1.8.2-0.20190607221125-9f4d16fe7c2d => github.com/nats-io/nats.go v1.8.1

replace github.com/asciiu/appa/lib => ../lib

require (
	github.com/asciiu/appa/api-graphql v0.0.5
	github.com/asciiu/appa/lib v0.0.0-00010101000000-000000000000
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.3.2
	github.com/google/uuid v1.1.1
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/lib/pq v1.3.0 // indirect
	github.com/lucas-clemente/quic-go v0.14.1 // indirect
	github.com/micro/examples v0.2.0
	github.com/micro/go-micro v1.18.0
	github.com/miekg/dns v1.1.27 // indirect
	github.com/nats-io/nats-server/v2 v2.1.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/stretchr/testify v1.4.0
	github.com/tmc/grpc-websocket-proxy v0.0.0-20200122045848-3419fae592fc // indirect
	go.uber.org/zap v1.13.0 // indirect
	golang.org/x/crypto v0.0.0-20200117160349-530e935923ad // indirect
	golang.org/x/lint v0.0.0-20191125180803-fdd1cda4f05f // indirect
	golang.org/x/net v0.0.0-20200114155413-6afb5195e5aa // indirect
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e // indirect
	golang.org/x/tools v0.0.0-20191216173652-a0e659d51361 // indirect
	google.golang.org/genproto v0.0.0-20191216164720-4f79533eabd1 // indirect
	google.golang.org/grpc v1.26.0 // indirect
	gopkg.in/libgit2/git2go.v27 v27.0.0-20190104134018-ecaeb7a21d47
)
