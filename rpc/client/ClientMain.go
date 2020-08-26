/*
@Time : 2020/8/26 10:27
@Author : ZhouHui2
@File : ClientMain
@Software: GoLand
*/
package client

import (
	"fmt"
	"log"
	"net/rpc"
)

func ClientMain() {
	client,err := rpc.Dial("tcp","localhost:1234")
	if err != nil {
		log.Fatal("Dial TCP error:", err)
	}
	var resp string
	err = client.Call("HelloService.Hello","hello",&resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}