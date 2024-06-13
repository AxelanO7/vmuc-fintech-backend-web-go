package usecase

import (
	"assyarif-backend-web-go/domain"
	"context"
	"fmt"
	"time"
)

type inUseCase struct {
	inRepository   domain.InRepository
	contextTimeout time.Duration
}

func NewInUseCase(in domain.InRepository, t time.Duration) domain.InUseCase {
	return &inUseCase{
		inRepository:   in,
		contextTimeout: t,
	}
}

func (c *inUseCase) ShowIns(ctx context.Context) ([]domain.In, error) {
	return c.inRepository.RetrieveIns()
}

func (c *inUseCase) ShowInLastNumber(ctx context.Context) (int, error) {
	var res []domain.In
	res, err := c.inRepository.RetrieveIns()
	if err != nil {
		return 0, err
	}

	lastNumber := 0
	for _, v := range res {
		fmt.Println(v.ID)
		if v.IdStuff > uint(lastNumber) {
			lastNumber = int(v.IdStuff)
		}
	}

	fmt.Println(lastNumber)
	return lastNumber, nil
}

func (c *inUseCase) AddIn(ctx context.Context, in domain.In) (domain.In, error) {
	return c.inRepository.CreateIn(in)
}

func (c *inUseCase) ShowInById(ctx context.Context, id string) (domain.In, error) {
	return c.inRepository.RetrieveInById(id)
}

func (c *inUseCase) EditInById(ctx context.Context, in domain.In) (domain.In, error) {
	return c.inRepository.UpdateInById(in)
}

func (c *inUseCase) DeleteInById(ctx context.Context, id string) error {
	return c.inRepository.RemoveInById(id)
}
