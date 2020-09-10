/*
@Time : 2020/9/10 15:42
@Author : ZhouHui2
@File : Stack_test.go
@Software: GoLand
*/
package LinkedListStack

import (
	"fmt"
	"log"
	"testing"
)

func TestInitNewStack(t *testing.T) {
	stack := InitNewStack()

	err := stack.Push(1)
	err = stack.Push(2)
	err = stack.Push(3)
	if err != nil {
		log.Fatal("push error")
	}
	result, err := stack.IsExist(1)
	fmt.Println(result)

}
