package application

import (
	"context"
	"fmt"
	"incrementer/storage"
)

type Application struct {
	Storage storage.Storage
}

func (app *Application) IncrementV1(ctx context.Context, i int) (int, error) {
	hasFirst, hasSecond, err := app.Storage.HasEither(ctx, i, i+1)
	if err != nil {
		return 0, fmt.Errorf("failed to find items: %w", err)
	}
	if hasFirst {
		return 0, fmt.Errorf("already done '%d'", i)
	}
	if hasSecond {
		return 0, fmt.Errorf("already done '%d'", i+1)
	}
	err = app.Storage.Insert(ctx, i)
	if err != nil {
		return 0, fmt.Errorf("failed to insert '%d': %w", i, err)
	}
	return i + 1, nil
}
