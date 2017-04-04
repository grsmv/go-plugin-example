package main

import (
	"context"
	"go-plugin-example/models"
)

func Weight() int {
	return 10
}

func Handler(ctx context.Context, data models.Data) (models.Data, error) {
	return models.Data{A: data.A + 100}, nil
}
