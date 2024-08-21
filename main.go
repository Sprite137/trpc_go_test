package main

import (
	"encoding/json"
	trpc "git.code.oa.com/trpc-go/trpc-go"
	thttp "git.code.oa.com/trpc-go/trpc-go/http"
	"io"
	"log"
	"net/http"
)

func main() {
	s := trpc.NewServer()
	// 路由注册 这里的HelloWorld是一个函数需要声明，用于实现这个路由注册的post请求的具体内容
	thttp.HandleFunc("/v1/hello", HelloWorld)
	// 服务注册
	thttp.RegisterNoProtocolService(s.Service("trpc.hello.stdhttp"))
	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}

// HelloWorld 实现handleFunc跳转的函数
func HelloWorld(w http.ResponseWriter, r *http.Request) error {
	// 获取请求报文中的数据
	msg, _ := io.ReadAll(r.Body)
	log.Printf("data is %s", msg)
	w.Header().Set("Content-type", "application/json")
	// 为响应报文设置 HTTP 状态码
	// w.WriteHeader(403)

	// 为响应报文设置 Body 这里的Data是一个结构体，需要额外声明
	rsp, _ := json.Marshal(&Data{Msg: "Hello, World!"})
	w.Write(rsp)

	return nil
}

// Data 声明结构体
type Data struct {
	Msg string
}
