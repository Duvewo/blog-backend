package cache

import (
	"github.com/go-redis/redis"
	"time"
)

type Cache struct {
	Client *redis.Client
}

func New() (*Cache, error) {
	r := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	if _, err := r.Ping().Result(); err != nil {
		return nil, err
	}

	return &Cache{Client: r}, nil

}

func (c *Cache) AddValue(key string, value interface{}, expires time.Duration) error {
	return c.Client.Set(key, value, expires).Err()
}

func (c *Cache) GetValue(key string) string {
	result, err := c.Client.Get(key).Result()

	if err != nil {
		panic(err)
	}

	return result

}

func (c *Cache) KeyExists(key string) bool {
	res, err := c.Client.Exists(key).Result()

	if err != nil {
		panic(err)
	}

	return res > 0

}

func (c *Cache) DeleteValue(key string) {
	c.Client.Del(key)
}