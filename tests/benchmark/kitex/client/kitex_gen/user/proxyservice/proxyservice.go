// Code generated by Kitex v0.7.1. DO NOT EDIT.

package proxyservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	user "github.com/kitex-contrib/codec-dubbo/tests/benchmark/kitex/client/kitex_gen/user"
)

func serviceInfo() *kitex.ServiceInfo {
	return proxyServiceServiceInfo
}

var proxyServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "ProxyService"
	handlerType := (*user.ProxyService)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetUser": kitex.NewMethodInfo(getUserHandler, newProxyServiceGetUserArgs, newProxyServiceGetUserResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "user",
		"ServiceFilePath": `api.thrift`,
		"IDLAnnotations": map[string][]string{
			"JavaClassName": []string{"org.apache.dubbo.UserProviderProxy"},
		},
	}

	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		KiteXGenVersion: "v0.7.1",
		Extra:           extra,
	}
	return svcInfo
}

func getUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.ProxyServiceGetUserArgs)
	realResult := result.(*user.ProxyServiceGetUserResult)
	success, err := handler.(user.ProxyService).GetUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newProxyServiceGetUserArgs() interface{} {
	return user.NewProxyServiceGetUserArgs()
}

func newProxyServiceGetUserResult() interface{} {
	return user.NewProxyServiceGetUserResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetUser(ctx context.Context, req *user.Request) (r *user.User, err error) {
	var _args user.ProxyServiceGetUserArgs
	_args.Req = req
	var _result user.ProxyServiceGetUserResult
	if err = p.c.Call(ctx, "GetUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
