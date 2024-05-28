package usecase

import (
	"assyarif-backend-web-go/domain"
	"context"
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
	return c.inRepository.RetrieveInLastNumber()
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
