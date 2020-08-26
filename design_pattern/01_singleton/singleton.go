/*
@Time : 2020/3/31 17:28
@Author : ZhouHui2
@File : singleton
@Software: GoLand
*/
package _1_singleton

import (
	"fmt"
	"sync"
)

/**
单例模式:
意图：保证一个类仅有一个实例，并提供一个访问它的全局访问点。
主要解决：一个全局使用的类频繁地创建与销毁。
何时使用：当您想控制实例数目，节省系统资源的时候。
如何解决：判断系统是否已经有这个单例，如果有则返回，如果没有则创建。
关键代码：构造函数是私有的。
*/

type singleton struct {
	data int
}

var sin *singleton

var once sync.Once

func GetSingleton() *singleton {
	once.Do(func() {
		fmt.Println("chushihua")
		sin = &singleton{12}
	})
	fmt.Println("实例对象的信息和地址", sin, &sin)
	return sin
}

func (t *singleton) Add() {
	t.data += 1
}
