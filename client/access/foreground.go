package access

import (
	accesspb "github.com/smile-im/microkit-client/proto/accesspb"
	"github.com/micro-kit/microkit/client"
	"google.golang.org/grpc"
)

var (
	svcName = "access"
)

// NewClient 创建客户端
func NewClient() (accessClient accesspb.AccessClient, err error) {
	c, err := client.NewDefaultClient(client.ServiceName(svcName))
	if err != nil {
		return
	}
	// 连接服务端
	err = c.Dial(func(cc *grpc.ClientConn) {
		accessClient = accesspb.NewAccessClient(cc)
	})
	return
}
