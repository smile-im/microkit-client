package auth

import (
	"github.com/micro-kit/microkit/client"
	authpb "github.com/smile-im/microkit-client/proto/authpb"
	"google.golang.org/grpc"
)

var (
	svcName = "auth"
)

// NewClient 创建客户端
func NewClient() (authClient authpb.AuthClient, err error) {
	c, err := client.NewDefaultClient(client.ServiceName(svcName))
	if err != nil {
		return
	}
	// 连接服务端
	err = c.Dial(func(cc *grpc.ClientConn) {
		authClient = authpb.NewAuthClient(cc)
	})
	return
}
