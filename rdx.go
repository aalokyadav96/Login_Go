package main 

import (
    "context"
    "github.com/redis/go-redis/v9"
    "fmt"
)

var ctx = context.Background()
var rdx = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func rdxGet(key string) string {
    val, err := rdx.Get(ctx, key).Result()
    if err == redis.Nil {
        fmt.Println("key does not exist")
    } else if err != nil {
        panic(err)
    } else {
//        fmt.Println(val)
		return val
    }
	return "error"
}

func rdxSet(key string, value string) bool {
    err := rdx.Set(ctx, key, value, 0).Err()
    if err != nil {
        panic(err)
		return false
    }
	return true
}
