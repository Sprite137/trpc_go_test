// Code generated by trpc-go/trpc-go-cmdline v2.6.2. DO NOT EDIT.
// source: helloworld.proto

package helloworld

import (
	"context"
	"errors"
	"fmt"

	_ "git.code.oa.com/trpc-go/trpc-go"
	"git.code.oa.com/trpc-go/trpc-go/client"
	"git.code.oa.com/trpc-go/trpc-go/codec"
	_ "git.code.oa.com/trpc-go/trpc-go/http"
	"git.code.oa.com/trpc-go/trpc-go/server"
)

// START ======================================= Server Service Definition ======================================= START

// GreeterRpcService defines service.
type GreeterRpcService interface {
	SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error)

	SayHi(ctx context.Context, req *HelloRequest) (*HelloReply, error)
}

func GreeterRpcService_SayHello_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &HelloRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(GreeterRpcService).SayHello(ctx, reqbody.(*HelloRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func GreeterRpcService_SayHi_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &HelloRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(GreeterRpcService).SayHi(ctx, reqbody.(*HelloRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// GreeterRpcServer_ServiceDesc descriptor for server.RegisterService.
var GreeterRpcServer_ServiceDesc = server.ServiceDesc{
	ServiceName: "trpc.test.helloworld.GreeterRpc",
	HandlerType: ((*GreeterRpcService)(nil)),
	Methods: []server.Method{
		{
			Name: "/trpc.test.helloworld.GreeterRpc/SayHello",
			Func: GreeterRpcService_SayHello_Handler,
		},
		{
			Name: "/trpc.test.helloworld.GreeterRpc/SayHi",
			Func: GreeterRpcService_SayHi_Handler,
		},
	},
}

// RegisterGreeterRpcService registers service.
func RegisterGreeterRpcService(s server.Service, svr GreeterRpcService) {
	if err := s.Register(&GreeterRpcServer_ServiceDesc, svr); err != nil {
		panic(fmt.Sprintf("GreeterRpc register error:%v", err))
	}
}

// GreeterHttpService defines service.
type GreeterHttpService interface {
	SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error)

	SayHi(ctx context.Context, req *HelloRequest) (*HelloReply, error)
}

func GreeterHttpService_SayHello_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &HelloRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(GreeterHttpService).SayHello(ctx, reqbody.(*HelloRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func GreeterHttpService_SayHi_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &HelloRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(GreeterHttpService).SayHi(ctx, reqbody.(*HelloRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// GreeterHttpServer_ServiceDesc descriptor for server.RegisterService.
var GreeterHttpServer_ServiceDesc = server.ServiceDesc{
	ServiceName: "trpc.test.helloworld.GreeterHttp",
	HandlerType: ((*GreeterHttpService)(nil)),
	Methods: []server.Method{
		{
			Name: "/trpc.test.helloworld.GreeterHttp/SayHello",
			Func: GreeterHttpService_SayHello_Handler,
		},
		{
			Name: "/trpc.test.helloworld.GreeterHttp/SayHi",
			Func: GreeterHttpService_SayHi_Handler,
		},
	},
}

// RegisterGreeterHttpService registers service.
func RegisterGreeterHttpService(s server.Service, svr GreeterHttpService) {
	if err := s.Register(&GreeterHttpServer_ServiceDesc, svr); err != nil {
		panic(fmt.Sprintf("GreeterHttp register error:%v", err))
	}
}

// START --------------------------------- Default Unimplemented Server Service --------------------------------- START

type UnimplementedGreeterRpc struct{}

func (s *UnimplementedGreeterRpc) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, errors.New("rpc SayHello of service GreeterRpc is not implemented")
}
func (s *UnimplementedGreeterRpc) SayHi(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, errors.New("rpc SayHi of service GreeterRpc is not implemented")
}

type UnimplementedGreeterHttp struct{}

func (s *UnimplementedGreeterHttp) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, errors.New("rpc SayHello of service GreeterHttp is not implemented")
}
func (s *UnimplementedGreeterHttp) SayHi(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, errors.New("rpc SayHi of service GreeterHttp is not implemented")
}

// END --------------------------------- Default Unimplemented Server Service --------------------------------- END

// END ======================================= Server Service Definition ======================================= END

// START ======================================= Client Service Definition ======================================= START

// GreeterRpcClientProxy defines service client proxy
type GreeterRpcClientProxy interface {
	SayHello(ctx context.Context, req *HelloRequest, opts ...client.Option) (rsp *HelloReply, err error)

	SayHi(ctx context.Context, req *HelloRequest, opts ...client.Option) (rsp *HelloReply, err error)
}

type GreeterRpcClientProxyImpl struct {
	client client.Client
	opts   []client.Option
}

var NewGreeterRpcClientProxy = func(opts ...client.Option) GreeterRpcClientProxy {
	return &GreeterRpcClientProxyImpl{client: client.DefaultClient, opts: opts}
}

func (c *GreeterRpcClientProxyImpl) SayHello(ctx context.Context, req *HelloRequest, opts ...client.Option) (*HelloReply, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/trpc.test.helloworld.GreeterRpc/SayHello")
	msg.WithCalleeServiceName(GreeterRpcServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("test")
	msg.WithCalleeServer("helloworld")
	msg.WithCalleeService("GreeterRpc")
	msg.WithCalleeMethod("SayHello")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &HelloReply{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *GreeterRpcClientProxyImpl) SayHi(ctx context.Context, req *HelloRequest, opts ...client.Option) (*HelloReply, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/trpc.test.helloworld.GreeterRpc/SayHi")
	msg.WithCalleeServiceName(GreeterRpcServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("test")
	msg.WithCalleeServer("helloworld")
	msg.WithCalleeService("GreeterRpc")
	msg.WithCalleeMethod("SayHi")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &HelloReply{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

// GreeterHttpClientProxy defines service client proxy
type GreeterHttpClientProxy interface {
	SayHello(ctx context.Context, req *HelloRequest, opts ...client.Option) (rsp *HelloReply, err error)

	SayHi(ctx context.Context, req *HelloRequest, opts ...client.Option) (rsp *HelloReply, err error)
}

type GreeterHttpClientProxyImpl struct {
	client client.Client
	opts   []client.Option
}

var NewGreeterHttpClientProxy = func(opts ...client.Option) GreeterHttpClientProxy {
	return &GreeterHttpClientProxyImpl{client: client.DefaultClient, opts: opts}
}

func (c *GreeterHttpClientProxyImpl) SayHello(ctx context.Context, req *HelloRequest, opts ...client.Option) (*HelloReply, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/trpc.test.helloworld.GreeterHttp/SayHello")
	msg.WithCalleeServiceName(GreeterHttpServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("test")
	msg.WithCalleeServer("helloworld")
	msg.WithCalleeService("GreeterHttp")
	msg.WithCalleeMethod("SayHello")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &HelloReply{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *GreeterHttpClientProxyImpl) SayHi(ctx context.Context, req *HelloRequest, opts ...client.Option) (*HelloReply, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/trpc.test.helloworld.GreeterHttp/SayHi")
	msg.WithCalleeServiceName(GreeterHttpServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("test")
	msg.WithCalleeServer("helloworld")
	msg.WithCalleeService("GreeterHttp")
	msg.WithCalleeMethod("SayHi")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &HelloReply{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

// END ======================================= Client Service Definition ======================================= END
