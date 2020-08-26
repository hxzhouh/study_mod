package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func Test1(t *testing.T) {

	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("loop i is - %d\n", i)
		}(i)
	}
	//fmt.Println("aaa")
	wg.Wait()
	fmt.Println("Hello, Welcome to Go")
}
