module github.com/smile-im/microkit-client/proto/authpb

go 1.15

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.4
	go.etcd.io/bbolt => github.com/coreos/bbolt v1.3.4
    google.golang.org/grpc => google.golang.org/grpc v1.26.0 // grpc对etcd依赖问题
)
