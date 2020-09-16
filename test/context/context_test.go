package context

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestWithDeadline(t *testing.T) {
	tick := time.NewTicker(1 * time.Second)
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(15*time.Second)) // 一秒后超时

	i := 0
	for {
		<-tick.C
		time, ok := ctx.Deadline()
		if ok {
			log.Println(time)
		}
		fmt.Println(time)
		if i == 5 {
			tick.Stop()
			break
		}
		i++
	}

}
