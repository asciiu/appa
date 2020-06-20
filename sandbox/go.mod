module github.com/asciiu/appa/sandbox

go 1.14

replace github.com/asciiu/appa/lib => ../lib

require (
	github.com/asciiu/appa/lib v0.0.0-00010101000000-000000000000
	github.com/fabioberger/coinbase-go v0.0.0-20160522011833-8328539b18ab // indirect
	github.com/jszwec/csvutil v1.3.0 // indirect
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/preichenberger/go-coinbasepro/v2 v2.0.5 // indirect
	github.com/sirupsen/logrus v1.6.0 // indirect
	github.com/stretchr/testify v1.5.1
	github.com/ybbus/jsonrpc v2.1.2+incompatible
)
