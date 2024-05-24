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

func (c *inUseCase) GetIns(ctx context.Context) ([]domain.In, error) {
	return c.inRepository.RetrieveIns()
}
