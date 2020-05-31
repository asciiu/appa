module github.com/asciiu/appa/api-graphql-rocket

go 1.14

replace github.com/asciiu/appa/lib => ../lib

require (
	github.com/99designs/gqlgen v0.11.2
	github.com/asciiu/appa/lib v0.0.0-00010101000000-000000000000
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/dgryski/trifles v0.0.0-20190318185328-a8d75aae118c
	github.com/go-chi/chi v4.0.1+incompatible
	github.com/go-pg/pg/v10 v10.0.0-alpha.0
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/google/uuid v1.1.1
	github.com/gorilla/mux v1.6.1
	github.com/gorilla/websocket v1.4.1
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/joho/godotenv v1.3.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/mitchellh/mapstructure v1.1.2 // indirect
	github.com/onsi/ginkgo v1.12.0 // indirect
	github.com/onsi/gomega v1.9.0 // indirect
	github.com/rs/cors v1.7.0
	github.com/segmentio/ksuid v1.0.2
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.5.1
	github.com/tinrab/graphql-realtime-chat v0.0.0-20200314084614-3e34df8561b0
	github.com/tinrab/retry v1.0.0
	github.com/vektah/gqlparser/v2 v2.0.1
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37
	golang.org/x/text v0.3.2 // indirect
)
