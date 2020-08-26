/*
@Time : 2020/3/31 17:30
@Author : ZhouHui2
@File : singleton_test.go
@Software: GoLand
*/
package _1_singleton

import (
	"fmt"
	"testing"
)

func TestSingleton(t *testing.T) {
	//sin1 := GetSingleton()
	//sin2 := GetSingleton()
	//
	//if sin1 != sin2 {
	//	t.Error("实例对象不一样")
	//}
	for i := 0; i < 100; i++ {
		sin := GetSingleton()
		sin.Add()
	}
	fmt.Println(GetSingleton().data)

}
