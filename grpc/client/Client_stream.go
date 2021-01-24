package client

import (
	"context"
	pb "github.com/hxzhouh/study_mod/grpc/protobuf"
	"google.golang.org/grpc"
	"io"
	"log"
	"strconv"
	"time"
)

var streamClient pb.StreamClient

func conversations() {
	//调用服务端的Conversations方法，获取流
	stream, err := streamClient.Conversations(context.Background())
	if err != nil {
		log.Fatalf("get conversations stream err: %v", err)
	}
	for n := 0; n < 10000; n++ {
		err := stream.Send(&pb.StreamRequest{Question: "stream client rpc " + strconv.Itoa(n)})
		if err != nil {
			log.Fatalf("stream request err: %v", err)
		}
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Conversations get stream err: %v", err)
		}
		// 打印返回值
		log.Println(res.Answer)
		time.Sleep(1*time.Second)
	}
	//最后关闭流
	err = stream.CloseSend()
	if err != nil {
		log.Fatalf("Conversations close stream err: %v", err)
	}
}
const Address string = "118.24.75.216:8000"


func client() {
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	streamClient = pb.NewStreamClient(conn)
	conversations()
}