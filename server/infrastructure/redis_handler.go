package infrastructure

import (
	"github.com/garyburd/redigo/redis"
)

type RedisHandler struct {
	connection redis.Conn
}

func NewRedisHandler() *RedisHandler {
	connection, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		return nil
	}
	redisHandler := RedisHandler {
		connection: connection,
	}
	return &redisHandler
} 

func (handler *RedisHandler) Set(key string, value string) error {
	_, err := handler.connection.Do("SET", key, value)
	if err != nil {
		return err
	}
	return nil
}

func (handler *RedisHandler) Get(key string) (string, error) {
	value, err := redis.String(handler.connection.Do("GET", key))
	if  err != nil {
		return value, err
	}
	return value, nil
}

func (handler *RedisHandler) RPush(key string, values []string) error {
	for _, value := range values {
		_, err := handler.connection.Do("RPUSH", key, value)
		if err != nil {
			return err
		}
	}
	return nil
}

func (handler *RedisHandler) LPop(key string, number int) ([]string, error) {
	result := []string{}
	for i := 0; i < number; i++ {
		value, err := redis.String(handler.connection.Do("LPOP", key))	
		if err != nil {
			return result, err
		}
		result = append(result, value)
	}
	return result, nil
}