package main

import (
	"context"
	"io"

	pb "trpc.cros.userApi"
)

type userApiImpl struct {
	pb.UnimplementedUserApi
}

// GetUserInfo 获取用户信息
func (s *userApiImpl) GetUserInfo(
	ctx context.Context,
	req *pb.Id,
) (*pb.UserInfo, error) {
	rsp := &pb.UserInfo{}
	return rsp, nil
}

// SaveUser 保存用户
func (s *userApiImpl) SaveUser(
	ctx context.Context,
	req *pb.UserParams,
) (*pb.Id, error) {
	rsp := &pb.Id{}
	return rsp, nil
}

func (s *userApiImpl) BatchSaveUser(
	stream pb.UserApi_BatchSaveUserServer,
) error {
	// 双端流式场景处理逻辑（仅供参考，请根据需要修改）
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		err = stream.Send(&pb.Id{})
		if err != nil {
			return err
		}
	}
}
