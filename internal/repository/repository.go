package repository

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type Repository interface {
	GetCount() (int, error)
	SetCount(count int) error
	CounterExists() error
}

type repository struct {
	client *redis.Client
}

func NewRepository(client *redis.Client) Repository {
	return &repository{
		client: client,
	}
}
func (r *repository) CounterExists() error {
	key := "request_counter"

	exists, err := r.client.Exists(context.Background(), key).Result()
	if err != nil {
		return fmt.Errorf("failed to check if key exists: %v", err)
	}

	if exists == 0 {
		err := r.client.Set(context.Background(), key, 0, 0).Err()
		if err != nil {
			return fmt.Errorf("failed to set key %s: %v", key, err)
		}
		fmt.Printf("Key %s created with initial value 0\n", key)
	}

	return nil
}

func (r *repository) GetCount() (int, error) {
	key := "request_counter"

	valueStr, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		return -1, err
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return -1, fmt.Errorf("failed to convert Redis value to integer: %v", err)
	}

	return value, nil
}
func (r *repository) SetCount(count int) error {
	key := "request_counter"

	err := r.client.Set(context.Background(), key, count, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to set Redis value: %v", err)
	}

	return nil
}
