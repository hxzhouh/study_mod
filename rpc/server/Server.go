/*
@Time : 2020/8/26 10:13
@Author : ZhouHui2
@File : Server
@Software: GoLand
*/
package server

import (
	"github.com/hxzhouh/study_mod/rpc/utils"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type SayHello struct {
}

type HelloServiceInterface = interface {
	Hello(req string, resp *string) error
}

func (t *SayHello) Hello(req string, resp *string) error {
	*resp = "Hello:" + req

	return nil
}

func ServerMain() {
	_ = RegisterHelloService(new(SayHello))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listen TCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(utils.HelloServiceName, svc)
}
