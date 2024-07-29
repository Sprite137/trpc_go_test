package main

import (
	"context"

	pb "git.woa.com/trpcprotocol/test/helloworld"
)

type greeterHttpImpl struct {
	pb.UnimplementedGreeterHttp
}

func (s *greeterHttpImpl) SayHello(
	ctx context.Context,
	req *pb.HelloRequest,
) (*pb.HelloReply, error) {
	rsp := &pb.HelloReply{}
	return rsp, nil
}

func (s *greeterHttpImpl) SayHi(
	ctx context.Context,
	req *pb.HelloRequest,
) (*pb.HelloReply, error) {
	rsp := &pb.HelloReply{}
	return rsp, nil
}
