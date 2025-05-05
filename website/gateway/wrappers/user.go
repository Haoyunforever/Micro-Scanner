package wrappers

import (
	"context"

	"github.com/afex/hystrix-go/hystrix"
	"go-micro.dev/v4/client"
)

// 基于Hystrix 的断路器模式的包装器，增强Go Micro客户端的容错能力。
// 它通过 Hystrix 提供的熔断机制，能够有效地防止服务调用的级联失败，同时在服务不可用时提供降级处理。
type userWrapper struct {
	client.Client
}

// 重写client.client接口方法
func (wrapper *userWrapper) Call(ctx context.Context, req client.Request, resp interface{}, opts ...client.CallOption) error {
	cmdName := req.Service() + "." + req.Endpoint() //生成命令名称，格式为服务点.端点名

	config := hystrix.CommandConfig{
		Timeout:                30000,
		RequestVolumeThreshold: 20,   // 熔断器请求阈值，默认20，意思是有20个请求才能进行错误百分比计算
		ErrorPercentThreshold:  50,   // 错误百分比，当错误超过百分比时，直接进行降级处理，直至熔断器再次 开启，默认50%
		SleepWindow:            5000, // 过多长时间，熔断器再次检测是否开启，单位毫秒ms（默认5秒）
	}
	hystrix.ConfigureCommand(cmdName, config)
	return hystrix.Do(cmdName, func() error {
		return wrapper.Client.Call(ctx, req, resp)
	}, func(err error) error {
		return err
	})
}

// NewUserWrapper 初始化Wrapper实例
func NewUserWrapper(c client.Client) client.Client {
	return &userWrapper{c}
}
