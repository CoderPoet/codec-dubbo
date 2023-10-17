// Code generated by Kitex v0.7.2. DO NOT EDIT.

package greetservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	hello "github.com/kitex-contrib/codec-dubbo/samples/helloworld/kitex/kitex_gen/hello"
)

func serviceInfo() *kitex.ServiceInfo {
	return greetServiceServiceInfo
}

var greetServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "GreetService"
	handlerType := (*hello.GreetService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Greet":           kitex.NewMethodInfo(greetHandler, newGreetServiceGreetArgs, newGreetServiceGreetResult, false),
		"GreetWithStruct": kitex.NewMethodInfo(greetWithStructHandler, newGreetServiceGreetWithStructArgs, newGreetServiceGreetWithStructResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "hello",
		"ServiceFilePath": `api.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		KiteXGenVersion: "v0.7.2",
		Extra:           extra,
	}
	return svcInfo
}

func greetHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*hello.GreetServiceGreetArgs)
	realResult := result.(*hello.GreetServiceGreetResult)
	success, err := handler.(hello.GreetService).Greet(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = &success
	return nil
}
func newGreetServiceGreetArgs() interface{} {
	return hello.NewGreetServiceGreetArgs()
}

func newGreetServiceGreetResult() interface{} {
	return hello.NewGreetServiceGreetResult()
}

func greetWithStructHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*hello.GreetServiceGreetWithStructArgs)
	realResult := result.(*hello.GreetServiceGreetWithStructResult)
	success, err := handler.(hello.GreetService).GreetWithStruct(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newGreetServiceGreetWithStructArgs() interface{} {
	return hello.NewGreetServiceGreetWithStructArgs()
}

func newGreetServiceGreetWithStructResult() interface{} {
	return hello.NewGreetServiceGreetWithStructResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Greet(ctx context.Context, req string) (r string, err error) {
	var _args hello.GreetServiceGreetArgs
	_args.Req = req
	var _result hello.GreetServiceGreetResult
	if err = p.c.Call(ctx, "Greet", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GreetWithStruct(ctx context.Context, req *hello.GreetRequest) (r *hello.GreetResponse, err error) {
	var _args hello.GreetServiceGreetWithStructArgs
	_args.Req = req
	var _result hello.GreetServiceGreetWithStructResult
	if err = p.c.Call(ctx, "GreetWithStruct", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
