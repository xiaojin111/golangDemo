package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile)
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Println(err)
	}
	log.Println(pong)

	rdb.Set(context.Background(), "nameq", "张三", 10*time.Second)
	name, _ := rdb.Get(context.Background(), "nameq").Result()
	log.Println(name)

}
