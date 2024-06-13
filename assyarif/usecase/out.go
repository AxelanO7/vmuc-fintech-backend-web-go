package usecase

import (
	"assyarif-backend-web-go/domain"
	"context"
	"fmt"
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
	var res []domain.Out
	res, err := c.outRepository.RetrieveOuts()
	if err != nil {
		return 0, err
	}

	lastNumber := 0
	for _, v := range res {
		fmt.Println(v.ID)
		if v.ID > uint(lastNumber) {
			lastNumber = int(v.ID)
		}

	}

	fmt.Println(lastNumber)
	return lastNumber, nil

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
