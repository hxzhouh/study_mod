package main

import "fmt"

func main() {
	c1 := make(chan bool, 1)
	c2 := make(chan bool, 1)
	for {
		select {
		case <-c1:
		case <-c2:
		default:
			fmt.Print("hello")

		}
	}

}
