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

func Test_greeterHttpImpl_SayHello(t *testing.T) {
	// 开始写mock逻辑
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	greeterHttpService := pb.NewMockGreeterHttpService(ctrl)
	var inorderClient []*gomock.Call
	// 预期行为
	m := greeterHttpService.EXPECT().SayHello(gomock.Any(), gomock.Any()).AnyTimes()
	m.DoAndReturn(func(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
		s := &greeterHttpImpl{}
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
			if rsp, err = greeterHttpService.SayHello(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("greeterHttpImpl.SayHello() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(rsp, tt.args.rsp) {
				t.Errorf("greeterHttpImpl.SayHello() rsp got = %v, want %v", rsp, tt.args.rsp)
			}
		})
	}
}

func Test_greeterHttpImpl_SayHi(t *testing.T) {
	// 开始写mock逻辑
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	greeterHttpService := pb.NewMockGreeterHttpService(ctrl)
	var inorderClient []*gomock.Call
	// 预期行为
	m := greeterHttpService.EXPECT().SayHi(gomock.Any(), gomock.Any()).AnyTimes()
	m.DoAndReturn(func(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
		s := &greeterHttpImpl{}
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
			if rsp, err = greeterHttpService.SayHi(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("greeterHttpImpl.SayHi() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(rsp, tt.args.rsp) {
				t.Errorf("greeterHttpImpl.SayHi() rsp got = %v, want %v", rsp, tt.args.rsp)
			}
		})
	}
}
