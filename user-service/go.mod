module github.com/asciiu/appa/user-service

go 1.12

require (
	github.com/armon/go-metrics v0.0.0-20160717043458-3df31a1ada83
	github.com/asciiu/appa v0.0.0
	github.com/go-log/log v0.1.0
	github.com/golang/protobuf v1.2.0
	github.com/google/uuid v1.0.0
	github.com/hashicorp/consul v1.3.0
	github.com/hashicorp/errwrap v1.0.0
	github.com/hashicorp/go-cleanhttp v0.0.0-20171218145408-d5fe4b57a186
	github.com/hashicorp/go-msgpack v0.0.0-20150518234257-fa3f63826f7c
	github.com/hashicorp/go-multierror v1.0.0
	github.com/hashicorp/go-rootcerts v0.0.0-20160503143440-6bb64b370b90
	github.com/hashicorp/go-sockaddr v0.0.0-20180320115054-6d291a969b86
	github.com/hashicorp/memberlist v0.1.0
	github.com/hashicorp/serf v0.0.0-20180809141758-19bbd39e421b
	github.com/json-iterator/go v1.1.5
	github.com/lib/pq v1.2.0
	github.com/micro/cli v0.0.0-20180830071301-8b9d33ec2f19
	github.com/micro/go-grpc v0.3.0
	github.com/micro/go-log v0.0.0-20170512141327-cbfa9447f9b6
	github.com/micro/go-micro v0.15.0
	github.com/micro/go-plugins v0.15.0
	github.com/micro/go-rcache v0.0.0-20180418165751-a581a57b5133
	github.com/micro/grpc-go v0.0.0-20180913204047-2c703400301b
	github.com/micro/h2c v1.0.0
	github.com/micro/kubernetes v0.1.0
	github.com/micro/mdns v0.0.0-20160929165650-cdf30746f9f7
	github.com/micro/util v0.0.0-20180417104657-4b7ed83e8520
	github.com/miekg/dns v1.0.15
	github.com/mitchellh/go-homedir v1.0.0
	github.com/mitchellh/hashstructure v1.0.0
	github.com/mitchellh/mapstructure v0.0.0-20181001021442-5a380f224700
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd
	github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742
	github.com/pkg/errors v0.8.0
	github.com/sean-/seed v0.0.0-20170313163322-e2103e2c3529
	golang.org/x/crypto v0.0.0-20181030102418-4d3f4d9ffa16
	golang.org/x/net v0.0.0-20181102091132-c10e9556a7bc
	golang.org/x/sys v0.0.0-20181031143558-9b800f95dbbc
	golang.org/x/text v0.3.0
	google.golang.org/genproto v0.0.0-20181101192439-c830210a61df
)

replace github.com/asciiu/appa => ../
