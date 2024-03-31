// redis_operations.go
package redisOperations

import (
	"assignment/sqlOperations"
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// RedisClient holds the Redis client instance
var RedisClient *redis.Client

// InitializeRedis initializes a connection to Redis
func InitializeRedis() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	fmt.Println("Connected to Redis database")
	return nil
}

// SetEmployee sets an employee record in Redis cache
func SetEmployee(key string, employee interface{}) error {
	jsonData, err := json.Marshal(employee)
	if err != nil {
		return err
	}

	err = RedisClient.Set(context.Background(), key, jsonData, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// GetEmployee retrieves an employee record from Redis cache
func GetEmployee(key string) (interface{}, error) {
	val, err := RedisClient.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}

	var employee interface{}
	err = json.Unmarshal([]byte(val), &employee)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

// DeleteEmployee deletes an employee record from Redis cache
func DeleteEmployee(key string) error {
	_, err := RedisClient.Del(context.Background(), key).Result()
	if err != nil {
		return err
	}
	return nil
}

// GetAllEmployeesFromRedis retrieves all employees from Redis cache
func GetAllEmployeesFromRedis() ([]sqlOperations.Employee, error) {
	var employees []sqlOperations.Employee

	keys, err := RedisClient.Keys(context.Background(), "employee:*").Result()
	if err != nil {
		return nil, err
	}

	for _, key := range keys {
		val, err := RedisClient.Get(context.Background(), key).Result()
		if err != nil {
			return nil, err
		}

		var employee sqlOperations.Employee
		err = json.Unmarshal([]byte(val), &employee)
		if err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	return employees, nil
}
