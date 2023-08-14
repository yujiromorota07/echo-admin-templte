package repository

import "context"

type TransactionManager interface {
	Do(ctx context.Context, runner func(ctx context.Context) (interface{}, error)) (interface{}, error)
}
