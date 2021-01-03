package auth

import (
	"github.com/micro-kit/microkit/client"
	authpb "github.com/smile-im/microkit-client/proto/authpb"
	"google.golang.org/grpc"
)

var (
	svcAdminName = "auth"
)

// NewAdminClient 创建管理端客户端
func NewAdminClient() (authAdminClient authpb.AdminAuthClient, err error) {
	c, err := client.NewDefaultClient(client.ServiceName(svcAdminName))
	if err != nil {
		return
	}
	// 连接服务端
	err = c.Dial(func(cc *grpc.ClientConn) {
		authAdminClient = authpb.NewAuthClient(cc)
	})
	return
}
