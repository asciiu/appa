module github.com/asciiu/appa/trade-engine

go 1.13.3

replace github.com/nats-io/nats.go v1.8.2-0.20190607221125-9f4d16fe7c2d => github.com/nats-io/nats.go v1.8.1

replace github.com/asciiu/appa/lib => ../lib

require (
	github.com/asciiu/appa/api-graphql v0.0.5
	github.com/asciiu/appa/lib v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.3.3
	github.com/google/uuid v1.1.1
	github.com/grpc-ecosystem/grpc-gateway v1.9.2 // indirect
	github.com/lib/pq v1.3.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.0.0
	github.com/stretchr/testify v1.4.0
)
