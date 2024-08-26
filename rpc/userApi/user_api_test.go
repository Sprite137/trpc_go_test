package main

import (
	"context"
	"io"
	"reflect"
	"testing"

	trpc "git.code.oa.com/trpc-go/trpc-go"
	_ "git.code.oa.com/trpc-go/trpc-go/http"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	pb "trpc.cros.userApi"
)

//go:generate go mod tidy
//go:generate mockgen -destination=stub/trpc.cros.userApi/userApi_mock.go -package=trpc_cros_userApi -self_package=trpc.cros.userApi --source=stub/trpc.cros.userApi/userApi.trpc.go

func Test_userApiImpl_GetUserInfo(t *testing.T) {
	// 开始写mock逻辑
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userApiService := pb.NewMockUserApiService(ctrl)
	var inorderClient []*gomock.Call
	// 预期行为
	m := userApiService.EXPECT().GetUserInfo(gomock.Any(), gomock.Any()).AnyTimes()
	m.DoAndReturn(func(ctx context.Context, req *pb.Id) (*pb.UserInfo, error) {
		s := &userApiImpl{}
		return s.GetUserInfo(ctx, req)
	})
	gomock.InOrder(inorderClient...)

	// 开始写单元测试逻辑
	type args struct {
		ctx context.Context
		req *pb.Id
		rsp *pb.UserInfo
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
			var rsp *pb.UserInfo
			var err error
			if rsp, err = userApiService.GetUserInfo(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("userApiImpl.GetUserInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(rsp, tt.args.rsp) {
				t.Errorf("userApiImpl.GetUserInfo() rsp got = %v, want %v", rsp, tt.args.rsp)
			}
		})
	}
}

func Test_userApiImpl_SaveUser(t *testing.T) {
	// 开始写mock逻辑
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userApiService := pb.NewMockUserApiService(ctrl)
	var inorderClient []*gomock.Call
	// 预期行为
	m := userApiService.EXPECT().SaveUser(gomock.Any(), gomock.Any()).AnyTimes()
	m.DoAndReturn(func(ctx context.Context, req *pb.UserParams) (*pb.Id, error) {
		s := &userApiImpl{}
		return s.SaveUser(ctx, req)
	})
	gomock.InOrder(inorderClient...)

	// 开始写单元测试逻辑
	type args struct {
		ctx context.Context
		req *pb.UserParams
		rsp *pb.Id
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
			var rsp *pb.Id
			var err error
			if rsp, err = userApiService.SaveUser(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("userApiImpl.SaveUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(rsp, tt.args.rsp) {
				t.Errorf("userApiImpl.SaveUser() rsp got = %v, want %v", rsp, tt.args.rsp)
			}
		})
	}
}

func Test_UserApi_BatchSaveUser(t *testing.T) {
	var userApiService = &userApiImpl{}

	// 开始写mock逻辑
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userApiClientProxy := pb.NewMockUserApiClientProxy(ctrl)
	batchSaveUserClient := pb.NewMockUserApi_BatchSaveUserClient(ctrl)
	batchSaveUserServer := pb.NewMockUserApi_BatchSaveUserServer(ctrl)
	inorderClient := make([]*gomock.Call, 0)
	inorderServer := make([]*gomock.Call, 0)

	// 预期行为
	m := userApiClientProxy.EXPECT().BatchSaveUser(gomock.Any(), gomock.Any()).AnyTimes()

	m.DoAndReturn(func(ctx context.Context, opts ...interface{}) (interface{}, error) {

		x := batchSaveUserClient.EXPECT().Send(gomock.Any()).AnyTimes()

		x.DoAndReturn(func(req interface{}) error {

			r, ok := req.(*pb.UserParams)
			if !ok {
				panic("invalid request")
			}

			s := batchSaveUserServer.EXPECT().Recv().Return(r, nil)
			inorderServer = append(inorderServer, s)
			return nil
		})

		k := batchSaveUserClient.EXPECT().CloseSend()

		k.DoAndReturn(func() error {

			batchSaveUserServer.EXPECT().Recv().Return(nil, io.EOF)

			err := userApiService.BatchSaveUser(batchSaveUserServer)
			if err != nil {
				return err
			}

			batchSaveUserClient.EXPECT().Recv().Return(nil, io.EOF)
			return nil
		})

		return batchSaveUserClient, nil
	})
	gomock.InOrder(inorderServer...)

	s := batchSaveUserServer.EXPECT().Send(gomock.Any()).AnyTimes()

	s.DoAndReturn(func(rsp interface{}) error {

		r, ok := rsp.(*pb.Id)
		if !ok {
			panic("invalid response")
		}

		c := batchSaveUserClient.EXPECT().Recv().Return(r, nil)
		inorderClient = append(inorderClient, c)
		return nil
	})
	gomock.InOrder(inorderClient...)

	// 开始写单元测试逻辑（仅供参考，请根据需要修改）
	stream, err := userApiClientProxy.BatchSaveUser(trpc.BackgroundContext())
	assert.Nil(t, err)
	assert.NotNil(t, stream)

	for i := 0; i < 5; i++ {

		req := &pb.UserParams{}

		// 输出每一次入参 (检查t.Logf输出，运行 `go test -v`)
		t.Logf("UserApi_BatchSaveUser req: %v", req)

		err := stream.Send(req)
		assert.Nil(t, err)
	}

	err = stream.CloseSend()
	assert.Nil(t, err)

	for {
		rsp, err := stream.Recv()
		if err == io.EOF {
			break
		}

		// 输出每一次返回 (检查t.Logf输出，运行 `go test -v`)
		t.Logf("UserApi_BatchSaveUser rsp: %v, err: %v", rsp, err)

		assert.Nil(t, err)
	}
}
