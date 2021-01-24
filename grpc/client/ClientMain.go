/*
@Time : 2020/8/28 16:57
@Author : ZhouHui2
@File : ClientMain
@Software: GoLand
*/
package client

import (
	"context"
	"github.com/hxzhouh/study_mod/grpc/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"time"
)

func ClientMain() {
	c, err := credentials.NewClientTLSFromFile("../conf/server.pem", "go-example-grpc")
	if err != nil {
		log.Fatalf("credentials.NewServerTLSFromFile err: %v", err)
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(c))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := protobuf.NewHelloClient(conn)

	// Contact the server and print out its response.
	name := "nosixtools"

	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		r, err := client.SayHello(ctx, &protobuf.HelloReq{Name: name})
		if err != nil {
			log.Fatal(err)

		} else {
			log.Printf("Hello: %s", r.Msg)
		}
		time.Sleep(time.Second)
	}
}
