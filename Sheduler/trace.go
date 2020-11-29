package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func trace1() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	fmt.Println("trace stop")
}
