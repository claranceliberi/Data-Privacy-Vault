package db

import "github.com/go-redis/redis"

var Client *redis.Client

func Init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,             // use default DB
	})

	// Ping Redis to check if the connection is working
	_, err := Client.Ping().Result()
	if err != nil {
		panic(err)
	}
}
