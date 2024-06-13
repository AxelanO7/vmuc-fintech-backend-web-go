package usecase

import (
	"assyarif-backend-web-go/domain"
	"context"
	"fmt"
	"time"
)

type outletUseCase struct {
	outletRepository domain.OutletRepository
	contextTimeout   time.Duration
}

func NewOutletUseCase(outlet domain.OutletRepository, t time.Duration) domain.OutletUseCase {
	return &outletUseCase{
		outletRepository: outlet,
		contextTimeout:   t,
	}
}

func (c *outletUseCase) FetchOutletByID(ctx context.Context, id uint) (*domain.Outlet, error) {
	res, err := c.outletRepository.RetrieveOutletByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *outletUseCase) FetchOutlets(ctx context.Context) ([]domain.Outlet, error) {
	res, err := c.outletRepository.RetrieveAllOutlet()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *outletUseCase) CreateOutlet(ctx context.Context, req *domain.Outlet) (*domain.Outlet, error) {
	res, err := c.outletRepository.CreateOutlet(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *outletUseCase) UpdateOutlet(ctx context.Context, req *domain.Outlet) (*domain.Outlet, error) {
	res, err := c.outletRepository.UpdateOutlet(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *outletUseCase) DeleteOutlet(ctx context.Context, id uint) error {
	err := c.outletRepository.DeleteOutlet(id)
	if err != nil {
		return err
	}
	return nil
}

func (c *outletUseCase) ShowOutletLastNumber(ctx context.Context) (int, error) {
	var res []domain.Outlet
	res, err := c.outletRepository.RetrieveAllOutlet()
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

func (c *outletUseCase) ShowOutletByIDUser(ctx context.Context, id uint) (*domain.Outlet, error) {
	res, err := c.outletRepository.ShowOutletByIDUser(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
