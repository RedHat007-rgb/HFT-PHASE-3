package redis

import "github.com/redis/go-redis/v9"




func ConnectToRedis() *redis.Client{
	url:="localhost:6379";
	client:=redis.NewClient(&redis.Options{
		Addr:url,
	})
	return client
}
