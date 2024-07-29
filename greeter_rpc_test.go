package main

import (
	"context"
	"reflect"
	"testing"

	_ "git.code.oa.com/trpc-go/trpc-go/http"

	"github.com/golang/mock/gomock"

	pb "git.woa.com/trpcprotocol/test/helloworld"
)

//go:generate go mod tidy
//go:generate mockgen -destination=stub/git.woa.com/trpcprotocol/test/helloworld/helloworld_mock.go -package=helloworld -self_package=git.woa.com/trpcprotocol/test/helloworld --source=stub/git.woa.com/trpcprotocol/test/helloworld/helloworld.trpc.go

func Test_greeterRpcImpl_SayHello(t *testing.T) {
	// 开始写mock逻辑
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	greeterRpcService := pb.NewMockGreeterRpcService(ctrl)
	var inorderClient []*gomock.Call
	// 预期行为
	m := greeterRpcService.EXPECT().SayHello(gomock.Any(), gomock.Any()).AnyTimes()
	m.DoAndReturn(func(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
		s := &greeterRpcImpl{}
		return s.SayHello(ctx, req)
	})
	gomock.InOrder(inorderClient...)

	// 开始写单元测试逻辑
	type args struct {
		ctx context.Context
		req *pb.HelloRequest
		rsp *pb.HelloReply
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rsp *pb.HelloReply
			var err error
			if rsp, err = greeterRpcService.SayHello(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("greeterRpcImpl.SayHello() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(rsp, tt.args.rsp) {
				t.Errorf("greeterRpcImpl.SayHello() rsp got = %v, want %v", rsp, tt.args.rsp)
			}
		})
	}
}

func Test_greeterRpcImpl_SayHi(t *testing.T) {
	// 开始写mock逻辑
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	greeterRpcService := pb.NewMockGreeterRpcService(ctrl)
	var inorderClient []*gomock.Call
	// 预期行为
	m := greeterRpcService.EXPECT().SayHi(gomock.Any(), gomock.Any()).AnyTimes()
	m.DoAndReturn(func(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
		s := &greeterRpcImpl{}
		return s.SayHi(ctx, req)
	})
	gomock.InOrder(inorderClient...)

	// 开始写单元测试逻辑
	type args struct {
		ctx context.Context
		req *pb.HelloRequest
		rsp *pb.HelloReply
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rsp *pb.HelloReply
			var err error
			if rsp, err = greeterRpcService.SayHi(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("greeterRpcImpl.SayHi() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(rsp, tt.args.rsp) {
				t.Errorf("greeterRpcImpl.SayHi() rsp got = %v, want %v", rsp, tt.args.rsp)
			}
		})
	}
}
