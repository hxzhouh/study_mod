package main

import (
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/hxzhouh/study_mod/etcd/internal"
	"log"
	"time"
)

func main() {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal("connect etcd error")
	}
	config := &internal.EtcdTools{Cli: cli, Timeout: 1 * time.Second}
	config.Put("wbb", "吴倍倍,我爱你")
	fmt.Println(config.Get("wbb"))
}
