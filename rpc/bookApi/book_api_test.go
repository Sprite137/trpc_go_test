package main

import (
	"context"
	"reflect"
	"testing"

	_ "git.code.oa.com/trpc-go/trpc-go/http"

	"github.com/golang/mock/gomock"

	pb "trpc.cros.bookApi"
)

//go:generate go mod tidy
//go:generate mockgen -destination=stub/trpc.cros.bookApi/bookApi_mock.go -package=trpc_cros_bookApi -self_package=trpc.cros.bookApi --source=stub/trpc.cros.bookApi/bookApi.trpc.go

func Test_bookApiImpl_GetBookInfo(t *testing.T) {
	// 开始写mock逻辑
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	bookApiService := pb.NewMockBookApiService(ctrl)
	var inorderClient []*gomock.Call
	// 预期行为
	m := bookApiService.EXPECT().GetBookInfo(gomock.Any(), gomock.Any()).AnyTimes()
	m.DoAndReturn(func(ctx context.Context, req *pb.BookParams) (*pb.Book, error) {
		s := &bookApiImpl{}
		return s.GetBookInfo(ctx, req)
	})
	gomock.InOrder(inorderClient...)

	// 开始写单元测试逻辑
	type args struct {
		ctx context.Context
		req *pb.BookParams
		rsp *pb.Book
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
			var rsp *pb.Book
			var err error
			if rsp, err = bookApiService.GetBookInfo(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("bookApiImpl.GetBookInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(rsp, tt.args.rsp) {
				t.Errorf("bookApiImpl.GetBookInfo() rsp got = %v, want %v", rsp, tt.args.rsp)
			}
		})
	}
}

func Test_bookApiImpl_SaveBook(t *testing.T) {
	// 开始写mock逻辑
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	bookApiService := pb.NewMockBookApiService(ctrl)
	var inorderClient []*gomock.Call
	// 预期行为
	m := bookApiService.EXPECT().SaveBook(gomock.Any(), gomock.Any()).AnyTimes()
	m.DoAndReturn(func(ctx context.Context, req *pb.Book) (*pb.BookId, error) {
		s := &bookApiImpl{}
		return s.SaveBook(ctx, req)
	})
	gomock.InOrder(inorderClient...)

	// 开始写单元测试逻辑
	type args struct {
		ctx context.Context
		req *pb.Book
		rsp *pb.BookId
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
			var rsp *pb.BookId
			var err error
			if rsp, err = bookApiService.SaveBook(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("bookApiImpl.SaveBook() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(rsp, tt.args.rsp) {
				t.Errorf("bookApiImpl.SaveBook() rsp got = %v, want %v", rsp, tt.args.rsp)
			}
		})
	}
}
