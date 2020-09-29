module github.com/asciiu/appa/sandbox-btc

go 1.14

replace github.com/asciiu/appa/lib => ../lib

require (
	github.com/Jeiwan/tinybit v0.0.0-20200123130446-78b29894463a
	github.com/fluidshare/plasma/lib v0.1.0 // indirect
	github.com/google/go-cmp v0.3.1
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.0.0
	golang.org/x/sys v0.0.0-20200923182605-d9f96fdee20d // indirect
)
