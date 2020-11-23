package add

import "fmt"

//go:nosplit
func add(a, b int) int // 汇编函数声明

func testAdd() {
	fmt.Println(add(10, 11))
}
