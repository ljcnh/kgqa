package redis

import (
	"github.com/gomodule/redigo/redis"
	"kgqa/config"
	"log"
	"strconv"
	"time"
)

// RedisPool redis的线程池
type RedisPool struct {
	Pool *redis.Pool
}

// PollInit redis 连接池初始化
func PollInit(config *config.RDConfig) RedisPool {
	address := config.Ip + ":" + strconv.Itoa(config.Port)
	defaultDB := config.DefaultDB
	passwd := config.Password
	return RedisPool{
		Pool: &redis.Pool{
			MaxIdle:     config.MaxIdle,
			MaxActive:   config.MaxActive,
			Wait:        true,
			IdleTimeout: time.Duration(config.IdleTimeout) * time.Second,
			Dial: func() (redis.Conn, error) {
				redis.DialDatabase(defaultDB)
				c, err := redis.Dial("tcp", address)
				if err != nil {
					log.Fatal(err)
					return nil, err
				}
				if _, err := c.Do("AUTH", passwd); err != nil {
					c.Close()
					return nil, err
				}
				return c, err
			},
		},
	}
}

func (redisPool *RedisPool) Close() error {
	if redisPool.Pool == nil {
		return nil
	}
	return redisPool.Pool.Close()
}

func (redisPool *RedisPool) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	conn := redisPool.Pool.Get()
	defer conn.Close()
	reply, err = conn.Do(commandName, args...)
	return
}
