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
)

func ClientMain() {
	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Accept error:", err)
	}
	var resp string
	client.Hello("wubeibei", &resp)

	fmt.Println(resp)
}
