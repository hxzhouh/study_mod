package Interceptor

import (
	"context"
	"fmt"
	"github.com/hxzhouh/study_mod/grpc/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type HelloServiceImpl struct{}

func (h HelloServiceImpl) SayHello(ctx context.Context, request *protobuf.HelloReq) (*protobuf.HelloResp, error) {
	result := &protobuf.HelloResp{
		Msg: fmt.Sprintf("%s%s", request.Name, "SayHello"),
	}
	return result, nil
}

func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("UnaryServerInterceptor before handling. Info: %+v", info)
	resp, err := handler(ctx, req)
	log.Printf("UnaryServerInterceptor after handling. resp: %+v", resp)
	return resp, err
}

// StreamServerInterceptor is a gRPC server-side interceptor that provides Prometheus monitoring for Streaming RPCs.
func StreamServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Printf("StreamServerInterceptor before handling. Info: %+v", info)
	err := handler(srv, ss)
	log.Printf("StreamServerInterceptor after handling. err: %v", err)
	return err
}

func ServerMain() {
	port := ":8000"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.StreamInterceptor(StreamServerInterceptor),
		grpc.UnaryInterceptor(UnaryServerInterceptor))
	protobuf.RegisterHelloServer(s, &HelloServiceImpl{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
