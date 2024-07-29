package main

import (
	"context"

	pb "git.woa.com/trpcprotocol/test/helloworld"
)

type greeterRpcImpl struct {
	pb.UnimplementedGreeterRpc
}

func (s *greeterRpcImpl) SayHello(
	ctx context.Context,
	req *pb.HelloRequest,
) (*pb.HelloReply, error) {
	rsp := &pb.HelloReply{}
	return rsp, nil
}

func (s *greeterRpcImpl) SayHi(
	ctx context.Context,
	req *pb.HelloRequest,
) (*pb.HelloReply, error) {
	rsp := &pb.HelloReply{}
	return rsp, nil
}
