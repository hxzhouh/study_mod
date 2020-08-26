/*
@Time : 2020/8/26 11:15
@Author : ZhouHui2
@File : client
@Software: GoLand
*/
package client

import (
	"github.com/hxzhouh/study_mod/rpc/utils"
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
	return t.Client.Call(utils.HelloServiceName+".Hello", req, resp)
}
