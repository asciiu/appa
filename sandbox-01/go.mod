module github.com/asciiu/appa/sandbox-01

go 1.14

replace github.com/asciiu/appa/lib => ../lib

replace github.com/blockcypher/libgrin => ../../../blockcypher/libgrin

require (
	github.com/asciiu/appa/lib v0.0.0-00010101000000-000000000000
	github.com/blockcypher/libgrin v2.0.0+incompatible
	github.com/btcsuite/btcd v0.20.1-beta
	github.com/ecies/go v1.0.1
	github.com/google/uuid v1.1.1
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/mitchellh/mapstructure v1.3.2 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/stretchr/testify v1.6.1
	github.com/tj/assert v0.0.3
	github.com/ybbus/jsonrpc v2.1.2+incompatible
)
