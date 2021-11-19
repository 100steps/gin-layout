package redis

import (
	"strconv"
	"time"

	"github.com/forseason/env"
	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func init() {
	maxIdle, err := strconv.Atoi(env.Get("REDIS_MAX_IDLE", "2"))
	if err != nil {
		panic(err)
	}
	maxActive, err := strconv.Atoi(env.Get("REDIS_MAX_ACTIVE", "3"))
	if err != nil {
		panic(err)
	}
	timeout, err := strconv.Atoi(env.Get("REDIS_IDLE_TIMEOUT", "240"))
	if err != nil {
		panic(err)
	}
	host, password := env.Get("REDIS_HOST", "127.0.0.1:6379"), env.Get("REDIS_PASSWORD", "")
	pool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: time.Duration(timeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", host)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := conn.Do("AUTH", password); err != nil {
					conn.Close()
					return nil, err
				}
			}
			return conn, err
		},
		TestOnBorrow: func(conn redis.Conn, time time.Time) error {
			_, err := conn.Do("PING")
			return err
		},
	}
}

// 用完连接后一定要记得将资源释放回连接池！
func GetConn() redis.Conn {
	return pool.Get()
}
