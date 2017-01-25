package cache

import (
	"fmt"

	"github.com/mediocregopher/radix.v2/redis"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisCache(host string, password string) (*RedisStore, error) {
	client, err := redis.Dial("tcp", host)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to redis: %v\n", err)
	}

	if password != "" {
		if err = client.Cmd("AUTH", password).Err; err != nil {
			client.Close()
			return nil, err
		}
	}

	return &RedisStore{client}, nil
}

func (c *RedisStore) Set(key string, value interface{}) error {
	err := c.client.Cmd("SET", value).Err
	if err != nil {
		return err
	}

	return nil
}

func (c *RedisStore) HashSet(key string, values []string) error {
	c.client.Cmd("HMSET", key, values)
	return nil
}

func (c *RedisStore) GetString(key string) (string, error) {
	val, err := c.client.Cmd("GET", key).Str()
	if err != nil {
		return "", err
	}

	return val, nil
}

func (c *RedisStore) GetInt(key string) (int, error) {
	val, err := c.client.Cmd("GET", key).Int()
	if err != nil {
		return 0, err
	}

	return val, nil
}

func (c *RedisStore) Delete(key string) error {
	err := c.client.Cmd("DEL", key).Err
	return err
}

func (c *RedisStore) Increment(key string, delta uint64) (uint64, error) {
	val, err := c.client.Cmd("INCRBY", key, delta).Int64()
	if err != nil {
		return 0, err
	}

	return uint64(val), nil
}

func (c *RedisStore) Decrement(key string, delta uint64) (newValue uint64, err error) {
	val, err := c.client.Cmd("DECRBY", key, delta).Int64()
	if err != nil {
		return 0, err
	}

	return uint64(val), nil
}
