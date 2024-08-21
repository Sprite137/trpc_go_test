package main

import (
	trpc "git.code.oa.com/trpc-go/trpc-go"
	thttp "git.code.oa.com/trpc-go/trpc-go/http"
	"hello/service"
	"log"
)

func main() {
	s := trpc.NewServer()

	service.InitService()

	// 路由注册 这里的HelloWorld是一个函数需要声明，用于实现这个路由注册的post请求的具体内容
	thttp.HandleFunc("/v1/hello", service.HelloWorld)
	// 服务注册
	thttp.RegisterNoProtocolService(s.Service("trpc.hello.stdhttp"))
	//thttp.RegisterNoProtocolService(s.Service("trpc.hello.http"))
	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}

	//pb.RegisterWemeetMinutesHttpService(s, service.SrvImp)
}
