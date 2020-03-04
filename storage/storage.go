package storage

import "context"

type Storage interface {
	HasEither(ctx context.Context, first, second int) (bool, bool, error)
	Insert(ctx context.Context, number int) error
}