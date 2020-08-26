package sync

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Encode(value string) []byte {
	valueByte := []byte(value)
	result := make([]byte, len(valueByte)/2)
	charpos, bufpos := 0, 0
	if len(valueByte)%2 == 1 {
		result[0] = valueByte[0] - 48
		charpos, bufpos = 1, 1
	}
	for ; charpos < len(valueByte); {
		result[bufpos] = (valueByte[charpos]-48)<<4 | (valueByte[charpos+1]-48)<<4
		charpos += 2
		bufpos += 1
	}
	return result

}

func TestEncode(t *testing.T) {
	fmt.Println(Encode("aaaa"))
}

var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)

func TestCond(t *testing.T) {
	for i := 0; i < 40; i++ {
		go func(x int) {
			cond.L.Lock()         //获取锁
			defer cond.L.Unlock() //释放锁
			cond.Wait()           //等待通知,阻塞当前goroutine
			fmt.Println(x)
			time.Sleep(time.Second * 1)

		}(i)
	}
	time.Sleep(time.Second * 1)
	fmt.Println("Signal...")
	cond.Signal() // 下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 1)
	cond.Signal() // 3秒之后 下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 3)
	cond.Broadcast() //3秒之后 下发广播给所有等待的goroutine
	fmt.Println("Broadcast...")
	time.Sleep(time.Second * 60)

}
