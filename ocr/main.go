package main

import (
	"flag"
	"fmt"
	"github.com/go-redis/redis/v7"
	"log"
)

func main() {
	var corpId int64
	flag.Int64Var(&corpId, "c", 21299, "corpId")
	flag.Parse()
	fmt.Println()

	client := redis.NewClient(&redis.Options{
		Addr:     "r-bp142699b6b62f14267.redis.rds.aliyuncs.com:6379",
		Password: "r-bp142699b6b62f14:BATy8M24xJ", // no password set
		DB:       0,                               // use default DB
	})
	//client := redis.NewClient(&redis.Options{
	//	Addr:     "192.168.1.52:6379",
	//	Password: "jh23jv78G2Ff9o0JHcw", // no password set
	//	DB:       0,                               // use default DB
	//})
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal("client error")
		return
	}
	result := client.Set(fmt.Sprintf("OCR_CORPID:%d", corpId), 1, -1)
	if result.Err() != nil {
		log.Fatal("set error", err)
	}else {
		log.Println("success,corpId=",corpId)
	}
}
