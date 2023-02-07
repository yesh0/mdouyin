// Code generated by Kitex v0.4.4. DO NOT EDIT.

package counterservice

import (
	"context"
	rpc "counter/kitex_gen/douyin/rpc"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Increment(ctx context.Context, req *rpc.CounterIncRequest, callOptions ...callopt.Option) (r *rpc.CounterNopResponse, err error)
	Fetch(ctx context.Context, req *rpc.CounterGetRequest, callOptions ...callopt.Option) (r *rpc.CounterGetResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kCounterServiceClient{
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

type kCounterServiceClient struct {
	*kClient
}

func (p *kCounterServiceClient) Increment(ctx context.Context, req *rpc.CounterIncRequest, callOptions ...callopt.Option) (r *rpc.CounterNopResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Increment(ctx, req)
}

func (p *kCounterServiceClient) Fetch(ctx context.Context, req *rpc.CounterGetRequest, callOptions ...callopt.Option) (r *rpc.CounterGetResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Fetch(ctx, req)
}
