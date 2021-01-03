package access

import (
	accesspb "github.com/smile-im/microkit-client/proto/accesspb"
	"github.com/micro-kit/microkit/client"
	"google.golang.org/grpc"
)

var (
	svcAdminName = "access"
)

// NewAdminClient 创建管理端客户端
func NewAdminClient() (accessAdminClient accesspb.AdminAccessClient, err error) {
	c, err := client.NewDefaultClient(client.ServiceName(svcAdminName))
	if err != nil {
		return
	}
	// 连接服务端
	err = c.Dial(func(cc *grpc.ClientConn) {
		accessAdminClient = accesspb.NewAccessClient(cc)
	})
	return
}
