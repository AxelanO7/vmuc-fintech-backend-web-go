package usecase

import (
	"assyarif-backend-web-go/domain"
	"context"
	"time"
)

type outUseCase struct {
	outRepository  domain.OutRepository
	contextTimeout time.Duration
}

func NewOutUseCase(out domain.OutRepository, t time.Duration) domain.OutUseCase {
	return &outUseCase{
		outRepository:  out,
		contextTimeout: t,
	}
}

func (c *outUseCase) GetOuts(ctx context.Context) ([]domain.Out, error) {
	return c.outRepository.RetrieveOuts()
}

func (c *outUseCase) ShowOutById(ctx context.Context, id string) (domain.Out, error) {
	return c.outRepository.RetrieveOutById(id)
}

func (c *outUseCase) ShowOutLastNumber(ctx context.Context) (int, error) {
	return c.outRepository.RetrieveOutLastNumber()
}

func (c *outUseCase) AddOut(ctx context.Context, out domain.Out) (domain.Out, error) {
	return c.outRepository.CreateOut(out)
}

func (c *outUseCase) EditOutById(ctx context.Context, out domain.Out) (domain.Out, error) {
	return c.outRepository.UpdateOutById(out)
}

func (c *outUseCase) DeleteOutById(ctx context.Context, id string) error {
	return c.outRepository.RemoveOutById(id)
}
