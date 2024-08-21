package service

import (
	_ "fmt"
)

func InitService() {
	SrvImp := &WebService{}
	SrvImp.name = "hello"
}
