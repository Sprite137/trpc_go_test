package main

import (
	"context"
	"git.code.oa.com/trpc-go/trpc-go/log"

	pb "trpc.cros.bookApi"
)

type bookApiImpl struct {
	pb.UnimplementedBookApi
}

// GetBookInfo 获取书籍信息
func (s *bookApiImpl) GetBookInfo(
	ctx context.Context,
	req *pb.BookParams,
) (*pb.Book, error) {
	rsp := &pb.Book{}
	log.Info("------GetBookInfo-----")
	return rsp, nil
}

// SaveBook 保存书籍
func (s *bookApiImpl) SaveBook(
	ctx context.Context,
	req *pb.Book,
) (*pb.BookId, error) {
	rsp := &pb.BookId{}
	return rsp, nil
}
