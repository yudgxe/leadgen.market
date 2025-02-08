package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var service HashService

type HashService interface {
	Save(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (bool, string, error)
}

type redisHashService struct {
	client *redis.Client
	ttl    int // in min
}

func (r redisHashService) Save(ctx context.Context, key, value string) error {
	if err := r.client.Set(ctx, key, value, time.Minute*time.Duration(r.ttl)).Err(); err != nil {
		return err
	}

	return nil
}

func (r redisHashService) Get(ctx context.Context, key string) (bool, string, error) {
	result, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return false, "", nil
	}
	if err != nil {
		return false, "", err
	}

	return true, result, nil
}

func SetupRedis(ctx context.Context, opt redis.Options, ttl int) error {
	client := redis.NewClient(&opt)

	if err := client.Ping(context.Background()).Err(); err != nil {
		return err
	}

	service = redisHashService{
		client: client,
		ttl:    ttl,
	}

	return nil
}

func GetHashService() HashService {
	return service
}

type mockHashService struct {}

func (s mockHashService) Save(ctx context.Context, key, value string) error {
	return nil
}

func (s mockHashService) Get(ctx context.Context, key string) (bool, string, error) {
	return false, "", nil
}

// SetupMock - сетапит мок в качествет кеша. Может использовать для тестов, либо в моментах, когда редис не доступен.
func SetupMock() {
	service = mockHashService{}
}