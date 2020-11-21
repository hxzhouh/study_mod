package main

import (
	"fmt"
	"sync/atomic"
)

const (
	mutexLocked      = 1 << iota // mutex is locked // 加锁状态
	mutexWoken                   //是否有唤醒的goroutine
	mutexStarving                // 从state字段中分出一个饥饿标记
	mutexWaiterShift = iota      // 等待的 goroutines

)

type Mutex struct {
	state uint32
	sema  uint32
}

func main() {

	fmt.Println(mutexLocked)
	fmt.Println(mutexWoken)
	fmt.Println(mutexStarving)
	fmt.Println(mutexWaiterShift)
	a := &Mutex{
		state: 0,
		sema:  0,
	}
	atomic.CompareAndSwapUint32(&a.state, 1, 0)
}
