package repository

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type Content struct {
	UUID     string
	Brand    string
	Model    string
	FuelType string
	Year     int32
}

type CarRepository struct {
	r *redis.Client
}

func NewCar(r *redis.Client) CarRepository {
	return CarRepository{
		r: r,
	}
}

func (c CarRepository) List(ctx context.Context) ([]Content, error) {
	keys, err := c.r.Keys(ctx, "*").Result()
	if err != nil {
		return nil, err
	}

	values, err := c.r.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, err
	}

	r := make([]Content, 0)
	for _, v := range values {
		var cc Content
		if err := json.Unmarshal([]byte(v.(string)), &cc); err != nil {
			return nil, err
		}

		r = append(r, cc)
	}

	return r, nil
}

func (c CarRepository) Save(ctx context.Context, content Content) (string, error) {
	content.UUID = uuid.NewString()

	contentMarshal, err := json.Marshal(content)
	if err != nil {
		return "", err
	}

	if err := c.r.Set(ctx, content.UUID, string(contentMarshal), 0).Err(); err != nil {
		return "", err
	}

	return content.UUID, nil
}
