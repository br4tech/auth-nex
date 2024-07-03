package port

import "context"

type IDatabase[T any] interface {
	FindOne(ctx context.Context, conds ...interface{}) (T, error)
	Create(ctx context.Context, value ...interface{}) (*any, error)
}
