package main

import (
	"fmt"
	"sync"
)

var i int64
var w sync.Mutex

func Inc() {
	w.Lock()
	i += 1
	defer w.Unlock()
}

var rw sync.RWMutex

func IncRW() {
	rw.Lock()
	i += 1
	defer rw.Unlock()
}
func main() {
	wg := sync.WaitGroup{}
	yawg := wg
	fmt.Println(wg, yawg)
}
