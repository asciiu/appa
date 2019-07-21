module github.com/asciiu/appa/story-service

go 1.12

require (
	github.com/asciiu/appa v0.0.0-00010101000000-000000000000
	github.com/asciiu/appa/api-graphql v0.0.0-20190721031036-120c5ff3b7dd
	github.com/asciiu/appa/user-service v0.0.0-20190721031036-120c5ff3b7dd
	github.com/golang/protobuf v1.3.2
	github.com/micro/go-micro v1.7.0
	github.com/stretchr/testify v1.3.0
	gopkg.in/libgit2/git2go.v27 v27.0.0-20190104134018-ecaeb7a21d47
)

replace github.com/asciiu/appa => ../

replace github.com/asciiu/api-graphql => ../api-graphql
