package main

import (
	"context"
	"go-plugin-example/models"
)

func Weight() int {
	return 20
}

func Handler(ctx context.Context, data models.Data) models.Data {
	return models.Data{A: data.A + 10}
}
