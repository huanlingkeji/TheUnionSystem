package main

import (
	"TheUnionSystem/github.com/go-redis/redis"
	"fmt"
)

func main() {
	client := redis.NewClient(&redis.Options{
		//Addr:     "172.26.147.228:6379",
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})

	//ping
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	//set key value
	err = client.Set("key2", "value2", 0).Err()
	if err != nil {
		panic(err)
	}

	//get key
	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	//get key2
	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

	//del key
	client.Del("key")

	client.Close()
}
