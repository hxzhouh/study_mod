package main

import "fmt"

//内存逃逸的情况
/**
1. 局部变量在外部有引用
2. 栈内存不够
3. 动态类型逃逸
*/

/**
栈上分配内存比在堆中分配内存有更高的效率
栈上分配的内存不需要GC处理
堆上分配的内存使用完毕会交给GC处理
逃逸分析目的是决定内分配地址是栈还是堆
逃逸分析在编译阶段完成
*/
type Student struct {
	Name string
	Age  int
}

func Slice() {
	s := make([]int, 1000, 10000) //栈空间不够,导致逃逸
	for index, _ := range s {
		s[index] = index
	}
}

func StudentRegister(name string, age int) *Student {
	s := new(Student) //局部变量s逃逸到堆
	s.Name = name
	s.Age = age
	return s
}

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	a := StudentRegister("Jim", 18)
	fmt.Print(a)
	s := "Escape"
	fmt.Println(s)
	f := Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("Fibonacci: %d\n", f())
	}
}
