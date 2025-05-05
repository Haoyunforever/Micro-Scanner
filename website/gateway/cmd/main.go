package main

import (
	"fmt"
	"github.com/Haoyunforever/Micro-Scanner/config"
	"github.com/Haoyunforever/Micro-Scanner/website/gateway/router"
	"github.com/Haoyunforever/Micro-Scanner/website/gateway/rpc"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"
	"time"

	"github.com/go-micro/plugins/v4/registry/etcd"
)

func main() {
	config.Init()
	rpc.InitRpc()
	etcReg := etcd.NewRegistry(registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)))

	//创建微服务实例，使用gin暴露http接口并注册到etcd中
	server := web.NewService(
		web.Name("ScannerService"),
		web.Version("latest"),
		web.Address("127.0.0.1:5000"),
		web.Handler(router.NewRouter()),
		web.Registry(etcReg),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Metadata(map[string]string{"protocol": "http"}),
	)

	_ = server.Init()
	_ = server.Run()
}
