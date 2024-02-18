package internal

import (
	"log"

	"github.com/redis/go-redis/v9"
)

type RedisServer struct {
	Conn *redis.Client
}

var RedisClient RedisServer

func InitializeRedis(host string, port string) {
	log.Printf("Connecting to Redis server at %s:%s", host, port)
	RedisClient.Conn = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: "",
		DB:       0,
	})
}

func GetRedisClient() *redis.Client {
	return RedisClient.Conn
}
