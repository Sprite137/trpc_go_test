// Package main 是由 trpc-go-cmdline v2.6.2 生成的客户端示例代码
// 本文件生成于 project/cmd/client 目录下
// 在 project 目录下执行 go run cmd/client/main.go 来运行本文件
// 注意：本文件并非必须存在，而仅为示例，用户应按需进行修改使用，如不需要，可直接删去
package main

import (
	_ "io"

	_ "git.code.oa.com/trpc-go/trpc-filter/debuglog"
	trpc "git.code.oa.com/trpc-go/trpc-go"
	"git.code.oa.com/trpc-go/trpc-go/client"
	"git.code.oa.com/trpc-go/trpc-go/log"
	bookApi "trpc.cros.bookApi"
	pb "trpc.cros.userApi"
)

func callUserApiGetUserInfo() {
	proxy := pb.NewUserApiClientProxy(
		client.WithTarget("ip://127.0.0.1:8000"),
		client.WithProtocol("trpc"),
	)
	ctx := trpc.BackgroundContext()
	// 一发一收 client 用法示例
	reply, err := proxy.GetUserInfo(ctx, &pb.Id{})
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	log.Debugf("simple  rpc   receive: %+v", reply)
}

func callBookApiGetBookInfo() {
	proxy := bookApi.NewBookApiClientProxy(
		client.WithTarget("ip://9.134.61.171:8001"),
		//client.WithTarget("ip://127.0.0.1:8001"),
		client.WithProtocol("trpc"),
	)
	ctx := trpc.BackgroundContext()
	// 一发一收 client 用法示例
	reply, err := proxy.GetBookInfo(ctx, &bookApi.BookParams{})
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	log.Debugf("simple  rpc   receive: %+v", reply)
}

func callUserApiSaveUser() {
	proxy := pb.NewUserApiClientProxy(
		client.WithTarget("ip://127.0.0.1:8000"),
		client.WithProtocol("trpc"),
	)
	ctx := trpc.BackgroundContext()
	// 一发一收 client 用法示例
	reply, err := proxy.SaveUser(ctx, &pb.UserParams{})
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	log.Debugf("simple  rpc   receive: %+v", reply)
}

func main() {
	// 仿照 trpc.NewServer 中的逻辑进行配置的加载
	cfg, err := trpc.LoadConfig(trpc.ServerConfigPath)
	if err != nil {
		panic("load config fail: " + err.Error())
	}
	trpc.SetGlobalConfig(cfg)
	if err := trpc.Setup(cfg); err != nil {
		panic("setup plugin fail: " + err.Error())
	}
	//callUserApiGetUserInfo()
	//callUserApiSaveUser()
	callBookApiGetBookInfo()
}
