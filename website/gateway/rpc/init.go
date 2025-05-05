package rpc

import (
	"github.com/Haoyunforever/Micro-Scanner/idl/pb"
	"github.com/Haoyunforever/Micro-Scanner/website/gateway/wrappers"
	"go-micro.dev/v4"
)

var UserService pb.UserService

func InitRpc() {
	userMicroService := micro.NewService(
		micro.Name("userService.client"),
		micro.WrapClient(wrappers.NewUserWrapper),
	)

	//用户服务调用实例
	userService := pb.NewUserService("rpcUserService.client", userMicroService.Client())

	UserService = userService
}
