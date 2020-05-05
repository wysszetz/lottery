package redis

import "github.com/go-redis/redis"

var Rdb *redis.Client

func initRdb() (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1",
		Password: "",
		DB:       0,
	})
	_, err = Rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
