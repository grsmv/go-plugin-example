package main

import (
	"context"
	"go-plugin-example/models"
	"time"
)

func Weight() int {
	return 10
}

func Handler(ctx context.Context, data models.Data) (models.Data, error) {
	time.Sleep(20 * time.Second)
	return models.Data{A: data.A + 100}, nil
}
