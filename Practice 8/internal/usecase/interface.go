package usecase

import "context"

type Usecase interface {
	CalculateSumConcurrent(ctx context.Context, n int) (result uint64, err error)
	CalculateSumLinear(ctx context.Context, n int) (result uint64, err error)
}
