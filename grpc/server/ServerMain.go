/*
@Time : 2020/8/28 16:50
@Author : ZhouHui2
@File : ServerMain
@Software: GoLand
*/
package server

import (
	"context"
	"fmt"
	"github.com/hxzhouh/study_mod/grpc/protobuf"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

type HelloServiceImpl struct{}

func (h HelloServiceImpl) SayHello(ctx context.Context, request *protobuf.HelloReq) (*protobuf.HelloResp, error) {
	result := &protobuf.HelloResp{
		Msg: fmt.Sprintf("%s%s", request.Name, "SayHello"),
	}
	return result, nil
}

func ServerMain() {
	port := ":8000"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protobuf.RegisterHelloServer(s, &HelloServiceImpl{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	mux:= http.NewServeMux()
	h2Handler :=h2c.NewHandler(mux,&http2.Server{})
	server :=http.Server{Addr: ":3999",Handler: h2Handler}
	server.ListenAndServe()
}
