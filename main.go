package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/go-redis/redis"
)

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func ping(client *redis.Client) error {
	pong, err := client.Ping().Result()
	if err != nil {
		return err
	}
	fmt.Println(pong, err)
	// Output: PONG <nil>

	return nil
}

func set(client *redis.Client, key string, value string, duration time.Duration) error {
	err := client.Set(key, value, duration).Err()
	if err != nil {
		return err
	}
	return nil
}

func get(client *redis.Client, key string) error {
	nameVal, err := client.Get(key).Result()
	if err != nil {
		return (err)
	}
	fmt.Println("name", nameVal)
	return nil
}

func main() {

	fmt.Println("Go Redis Tutorial")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	//set(client, "name", "risa", 0)
	//get(client, "name")

	pipe := client.Pipeline()
	pipe.Set("pipename2", "piperisa", 0)
	output := pipe.Get("pipename2")
	_, err := pipe.Exec()
	if err != nil {
		panic(err)
	} else {
		fmt.Println(reflect.TypeOf(output))
	}

}
