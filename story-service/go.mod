module github.com/asciiu/appa/story-service

go 1.12

replace github.com/nats-io/nats.go v1.8.2-0.20190607221125-9f4d16fe7c2d => github.com/nats-io/nats.go v1.8.1

require (
	github.com/asciiu/appa/api-graphql v0.0.5
	github.com/asciiu/appa/common v0.0.1
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/golang/protobuf v1.3.2
	github.com/micro/examples v0.2.1-0.20190728152459-704a6ae0074d // indirect
	github.com/micro/go-micro v1.8.0
	github.com/stretchr/testify v1.3.0
	gopkg.in/libgit2/git2go.v27 v27.0.0-20190104134018-ecaeb7a21d47 // indirect
)
