package usecase

import (
	"assyarif-backend-web-go/domain"
	"context"
	"time"
)

type opnameUseCase struct {
	opnameRepository domain.OpnameRepository
	contextTimeout   time.Duration
}

func NewOpnameUseCase(opname domain.OpnameRepository, t time.Duration) domain.OpnameUseCase {
	return &opnameUseCase{
		opnameRepository: opname,
		contextTimeout:   t,
	}
}

func (c *opnameUseCase) AddOpname(ctx context.Context, opname *domain.Opname) error {
	err := c.opnameRepository.CreateOpname(opname)
	if err != nil {
		return err
	}
	return nil
}

func (c *opnameUseCase) FetchOpnameByID(ctx context.Context, id uint) (*domain.Opname, error) {
	res, err := c.opnameRepository.RetrieveOpnameByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *opnameUseCase) FetchOpnames(ctx context.Context) ([]domain.Opname, error) {
	res, err := c.opnameRepository.RetrieveAllOpname()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *opnameUseCase) FetchOpnameByDate(ctx context.Context, startDate, endDate string) (*domain.ResByDate, error) {
	in, out, rtr, err := c.opnameRepository.RetriveByStartDateEndDate(startDate, endDate)
	if err != nil {
		return nil, err
	}
	return &domain.ResByDate{
		In:  in,
		Out: out,
		Rtr: rtr,
	}, nil
}
