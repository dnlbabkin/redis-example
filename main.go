package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Person struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Occupation string `json:"occupation"`
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(ping)

	err = client.Set(context.Background(), "name", "Danil", 0).Err()
	if err != nil {
		fmt.Printf("Failed to save value in redis %s", err.Error())
		return
	}

	val, err := client.Get(context.Background(), "name").Result()
	if err != nil {
		fmt.Printf("Cant show value from redis %s", err.Error())
		return
	}
	fmt.Println(val)

	jsonToStrin, err := json.Marshal(Person{
		Name:       "Danil",
		Age:        21,
		Occupation: "Chelik)",
	})
	if err != nil {
		fmt.Printf("Cant marshal json %s", err.Error())
	}

	err = client.Set(context.Background(), "person", jsonToStrin, 0).Err()
	if err != nil {
		fmt.Printf("Failed to save value in redis %s", err.Error())
		return
	}

	person, err := client.Get(context.Background(), "person").Result()
	if err != nil {
		fmt.Printf("Cant show value from redis %s", err.Error())
		return
	}
	fmt.Println(person)

}
