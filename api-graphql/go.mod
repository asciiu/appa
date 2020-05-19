module github.com/asciiu/appa/api-graphql

go 1.14

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.6.0

replace github.com/nats-io/nats.go v1.8.2-0.20190607221125-9f4d16fe7c2d => github.com/nats-io/nats.go v1.8.1

replace github.com/asciiu/appa/lib => ../lib

require (
	github.com/99designs/gqlgen v0.11.3
	github.com/asciiu/appa/api-graphql-rocket v0.0.0-20200311025607-409459594f2c // indirect
	github.com/asciiu/appa/lib v0.0.0-20200309030253-e7d07fc5fae3
	github.com/asciiu/appa/story-service v0.0.5
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/google/uuid v1.1.1
	github.com/joho/godotenv v1.3.0
	github.com/micro/examples v0.2.0
	github.com/micro/go-micro v1.9.1
	github.com/rs/cors v1.7.0
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.4.0
	github.com/vektah/gqlparser v1.2.0
	github.com/vektah/gqlparser/v2 v2.0.1
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550
)
