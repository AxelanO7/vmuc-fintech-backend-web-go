package usecase

import (
	"assyarif-backend-web-go/domain"
	"context"
	"time"
)

type outUseCase struct {
	outRepository   domain.OutRepository
	contextTimeout time.Duration
}

func NewOutUseCase(out domain.OutRepository, t time.Duration) domain.OutUseCase {
	return &outUseCase{
		outRepository: out,
		contextTimeout: t,
	}
}

func (c *outUseCase) GetOuts(ctx context.Context) ([]domain.Out, error) {
	return c.outRepository.RetrieveOuts()
}
