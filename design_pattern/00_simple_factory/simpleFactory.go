/*
@Time : 2020/3/31 15:40
@Author : ZhouHui2
@File : simpleFactory
@Software: GoLand
*/
package _0_simple_factory

import "fmt"

type Duck interface {
	Say(jiujiu string) string
}

func NewDuck(t int) Duck {
	if t == 1 {
		return &yellowDuck{}
	} else {
		return &blackDuck{}
	}
}

type yellowDuck struct{}

func (*yellowDuck) Say(jiujiu string) string {
	return fmt.Sprintf("YellowDuck say %s", jiujiu)
}

type blackDuck struct{}

func (*blackDuck) Say(jiujiu string) string {
	return fmt.Sprintf("BlackDuck say %s", jiujiu)
}
