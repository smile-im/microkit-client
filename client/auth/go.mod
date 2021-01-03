module github.com/smile-im/microkit-client/client/auth

go 1.15

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.4
	github.com/micro-kit/micro-common => ../../../../micro-kit/micro-common
	github.com/micro-kit/microkit => ../../../../micro-kit/microkit
	github.com/smile-im/microkit-client/proto/authpb => ../../proto/authpb
	go.etcd.io/bbolt => github.com/coreos/bbolt v1.3.4
	google.golang.org/grpc => google.golang.org/grpc v1.26.0 // grpc对etcd依赖问题
)

require (
	github.com/micro-kit/microkit v0.0.0-00010101000000-000000000000
	github.com/smile-im/microkit-client/proto/authpb v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.28.1
)
