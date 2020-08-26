/*
@Time : 2020/8/26 10:27
@Author : ZhouHui2
@File : ClientMain
@Software: GoLand
*/
package client

import (
	"fmt"
	"github.com/hxzhouh/study_mod/rpc/server"
	"log"
	"net/rpc"
)

func ClientMain() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Dial TCP error:", err)
	}
	var resp string
	funcName := fmt.Sprintf("%s.%s", server.HelloServiceName, "Hello")
	err = client.Call(funcName, "wubeibei", &resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
