/*
@Time : 2020/8/26 11:15
@Author : ZhouHui2
@File : client
@Software: GoLand
*/
package client

import (
	"github.com/hxzhouh/study_mod/rpc/utils"
	"net/rpc"
)

type HelloServiceClient struct {
	*rpc.Client
}

var _HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}
func (t *HelloServiceClient) Hello(req string, resp *string) error {
	return t.Client.Call(utils.HelloServiceName+".Hello", req, resp)
}
