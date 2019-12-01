package base

import (
	"TheUnionSystem/github.com/go-redis/redis"
	"time"
)

var redisCli *redis.Client

func Init(url string) {
	opts, e := redis.ParseURL(url)
	if e != nil {
		panic(e)
	}

	opts.DialTimeout = 10 * time.Second
	opts.ReadTimeout = 5 * time.Second
	opts.WriteTimeout = 5 * time.Second
	opts.PoolSize = 10
	opts.PoolTimeout = 30 * time.Second
	opts.IdleTimeout = 500 * time.Millisecond
	opts.IdleCheckFrequency = 500 * time.Millisecond
	redisCli = redis.NewClient(opts)

	_, err := redisCli.Ping().Result()
	if nil != err {
		panic("redis server not connect")
	}
}
