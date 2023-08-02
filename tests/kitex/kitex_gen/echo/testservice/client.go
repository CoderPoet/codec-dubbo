package testservice

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	hessian2 "github.com/kitex-contrib/codec-hessian2/pkg"
	"github.com/kitex-contrib/codec-hessian2/tests/kitex/kitex_gen/echo"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	EchoInt(ctx context.Context, req int32, callOptions ...callopt.Option) (r int32, err error)
	Echo(ctx context.Context, req *echo.EchoRequest, callOptions ...callopt.Option) (r *echo.EchoResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))
	options = append(options, client.WithCodec(hessian2.NewHessian2Codec()))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kTestServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kTestServiceClient struct {
	*kClient
}

func (p *kTestServiceClient) EchoInt(ctx context.Context, req int32, callOptions ...callopt.Option) (r int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoInt(ctx, req)
}

func (p *kTestServiceClient) Echo(ctx context.Context, req *echo.EchoRequest, callOptions ...callopt.Option) (r *echo.EchoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Echo(ctx, req)
}
