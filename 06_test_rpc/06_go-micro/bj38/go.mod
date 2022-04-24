module bj38

go 1.15

require (
	github.com/asim/go-micro/plugins/server/grpc/v3 v3.7.0
	github.com/asim/go-micro/plugins/transport/grpc/v3 v3.7.0
	github.com/asim/go-micro/v3 v3.7.1
	github.com/micro/micro/v3 v3.10.1
	github.com/prometheus/procfs v0.0.5 // indirect
	google.golang.org/protobuf v1.28.0
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.45.0
