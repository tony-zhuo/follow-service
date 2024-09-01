package model

import "context"

type WorkerInterface interface {
	ID(ctx context.Context) string
	Exec(ctx context.Context) error
	Topic(ctx context.Context) string
	Group(ctx context.Context) string
	Health(ctx context.Context) bool
}
