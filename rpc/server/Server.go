/*
@Time : 2020/8/26 10:13
@Author : ZhouHui2
@File : Server
@Software: GoLand
*/
package server

import (
	"log"
	"net"
	"net/rpc"
)

type SayHello struct {
}

func (t *SayHello) Hello(req string, resp *string) error {
	*resp = "Hello:" + req
	return nil
}

func ServerMain() {
	rpc.RegisterName("HelloService", new(SayHello))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listen TCP error:", err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}
	rpc.ServeConn(conn)
}
