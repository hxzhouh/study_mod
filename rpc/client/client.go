/*
@Time : 2020/8/26 11:15
@Author : ZhouHui2
@File : client
@Software: GoLand
*/
package client

import (
	"fmt"
	"github.com/hxzhouh/study_mod/rpc/utils"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloServiceClient struct {
	*rpc.Client
}

var _HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(c))
	return &HelloServiceClient{Client: client}, nil
}
func (t *HelloServiceClient) Hello(req string, resp *string) error {
	doClientWork(t.Client)
	return nil
	//return t.Client.Call(utils.HelloServiceName+".Hello", req, resp)
}

func doClientWork(client *rpc.Client) {
	helloCall := client.Go(utils.HelloServiceName+".Hello", "wubeibei", new(string), nil)
	// do some Things

	helloCall = <-helloCall.Done //这里会阻塞等待返回
	if err := helloCall.Error; err != nil {
		log.Fatal(err)
	}
	args := helloCall.Args.(string)
	reply := helloCall.Reply.(*string)
	fmt.Println(args, *reply)
}
