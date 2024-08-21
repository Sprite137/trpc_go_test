package service

import (
	"encoding/json"
	"hello/entity"
	"io"
	"log"
	"net/http"
)

type WebService struct {
	name string
}

func HelloWorld(w http.ResponseWriter, r *http.Request) error {
	// 获取请求报文中的数据
	msg, _ := io.ReadAll(r.Body)
	log.Printf("data is %s", msg)
	w.Header().Set("Content-type", "application/json")

	// 为响应报文设置 Body 这里的Data是一个结构体，需要额外声明
	rsp, _ := json.Marshal(&entity.Data{Msg: string(msg)})
	w.Write(rsp)

	return nil

}
