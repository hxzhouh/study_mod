package main

import (
	"fmt"
	"github.com/philchia/agollo/v4"
	"sync"
)

func getagolloConf(addId, cluster, cacheDir, mateAddr string, nameSpace []string) *agollo.Conf {
	return &agollo.Conf{
		AppID:          addId,
		Cluster:        cluster,
		NameSpaceNames: nameSpace,
		CacheDir:       cacheDir,
		MetaAddr:       mateAddr,
	}
}
func main() {
	agollo.Start(getagolloConf("fabio-timtraceloginsvr", "dev", ".", "192.168.1.65:8080", []string{"application"}))
	agollo.OnUpdate(func(event *agollo.ChangeEvent) {
		// do your business logic to handle config update
		fmt.Println(event)
		//switch event.Namespace {
		//case "application":
		//	handler(event.Changes)
		//}
	})
	val := agollo.GetString("cronValue")
	fmt.Println(val)
	agollo.OnUpdate(func(event *agollo.ChangeEvent) {
		switch event.Namespace {
		case "application":
			handler(event.Changes)
		}
	})
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}

func handler(v map[string]*agollo.Change) {
	for k, v := range v {
		fmt.Println("chanage key:", k)
		fmt.Println("change value:", v)
	}
}
