module github.com/asciiu/appa/trade-engine

go 1.13.3

replace github.com/nats-io/nats.go v1.8.2-0.20190607221125-9f4d16fe7c2d => github.com/nats-io/nats.go v1.8.1

require (
	github.com/asciiu/appa/api-graphql v0.0.5
	github.com/asciiu/appa/lib v0.0.0-20191022013234-7fdd8e2731b3
	github.com/golang/protobuf v1.3.2
	github.com/google/uuid v1.1.1
	github.com/lib/pq v1.2.0
	github.com/micro/examples v0.2.0
	github.com/micro/go-micro v1.9.1
	github.com/miekg/dns v1.1.17 // indirect
	github.com/stretchr/testify v1.3.0
	google.golang.org/genproto v0.0.0-20190819201941-24fa4b261c55 // indirect
	google.golang.org/grpc v1.23.0 // indirect
)
