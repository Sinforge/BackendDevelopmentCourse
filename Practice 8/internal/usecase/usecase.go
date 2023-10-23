package usecase

import (
	"context"
	"sync"
	"time"
)

type UserUseCase struct {
}

func New() Usecase {
	return &UserUseCase{}
}

func (uc *UserUseCase) CalculateFactorialConcurrent(_ context.Context, n int) (result uint64, err error) {
	result = 1
	wg := &sync.WaitGroup{}
	for i := 1; i <= n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			time.Sleep(time.Second)
			result *= uint64(i)
		}(i)
	}

	wg.Wait()
	return result, err
}
func (uc *UserUseCase) CalculateFactorialLinear(_ context.Context, n int) (result uint64, err error) {
	result = 1

	for i := 1; i <= n; i++ {
		time.Sleep(time.Second)
		result *= uint64(i)
	}

	return result, err
}
