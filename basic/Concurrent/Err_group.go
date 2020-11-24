package Concurrent

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func TestErrGroup() {

	var g errgroup.Group
	var result = make([]error, 3)
	g.Go(func() error {
		time.Sleep(5 * time.Second)
		fmt.Println("exec #1")
		result[0] = nil // 保存成功或者失败的结果
		return nil
	})
	g.Go(func() error {
		time.Sleep(5 * time.Second)
		fmt.Println("exec #2")
		result[1] = errors.New("你是傻逼吗？") // 保存成功或者失败的结果
		return nil
	})

	g.Go(func() error {
		time.Sleep(5 * time.Second)
		fmt.Println("exec #3")
		result[2] = nil // 保存成功或者失败的结果
		return nil
	})
	if err := g.Wait(); err != nil {
		fmt.Printf("Successfully exec all. result: %v\n", result)
	} else {
		fmt.Printf("failed: %v\n", result)
	}

}
