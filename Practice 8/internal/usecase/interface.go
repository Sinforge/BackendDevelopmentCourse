package usecase

import "context"

type Usecase interface {
	CalculateFactorialConcurrent(ctx context.Context, n int) (result uint64, err error)
	CalculateFactorialLinear(ctx context.Context, n int) (result uint64, err error)
}
