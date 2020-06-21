module github.com/asciiu/appa/sandbox-01

go 1.14

replace github.com/asciiu/appa/lib => ../lib

require (
	github.com/asciiu/appa/lib v0.0.0-00010101000000-000000000000
	github.com/ecies/go v1.0.1
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/sirupsen/logrus v1.6.0
	github.com/stretchr/testify v1.5.1
	github.com/ybbus/jsonrpc v2.1.2+incompatible
)
