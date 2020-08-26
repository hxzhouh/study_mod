package sync

import (
	"fmt"
	"sync"
	"testing"
)

var printPid sync.Once

func TestOnce(t *testing.T) {
	for i:=0;i<10;i++ {
		printPid.Do(func() {
			fmt.Println("hello,world")
		})
	}

}
