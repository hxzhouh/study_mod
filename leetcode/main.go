package main

import "fmt"

type humen interface {
	Run() string
	Laugh()
	man
}

type man struct {
	name string
	age  string
}

func (t *man) Run() string {
	return "man running"

}
func (t *man) Laugh() {

}

type woman struct {
	eye string
}

func (t *woman) Run() string {
	return "woman running"
}

func (t *woman) Laugh() {

}

type student struct {
	info humen
}

func main() {
	s := &student{info: &man{}}
	fmt.Println(s.info.Run())

	//fmt.Println(runtime.GOOS)

}
