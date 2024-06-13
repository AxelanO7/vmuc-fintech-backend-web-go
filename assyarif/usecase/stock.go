package usecase

import (
	"assyarif-backend-web-go/domain"
	"context"
	"time"
)

type stockUseCase struct {
	stockRepository domain.StockRepository
	contextTimeout  time.Duration
}

func NewStockUseCase(stock domain.StockRepository, t time.Duration) domain.StockUseCase {
	return &stockUseCase{
		stockRepository: stock,
		contextTimeout:  t,
	}
}

func (c *stockUseCase) FetchStockByID(ctx context.Context, id uint) (*domain.Stock, error) {
	res, err := c.stockRepository.RetrieveStockByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *stockUseCase) FetchStocks(ctx context.Context) ([]domain.Stock, error) {
	res, err := c.stockRepository.RetrieveAllStock()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *stockUseCase) CreateStock(ctx context.Context, req *domain.Stock) (*domain.Stock, error) {
	res, err := c.stockRepository.CreateStock(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *stockUseCase) UpdateStock(ctx context.Context, req *domain.Stock) (*domain.Stock, error) {
	res, err := c.stockRepository.UpdateStock(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *stockUseCase) DeleteStock(ctx context.Context, id uint) error {
	err := c.stockRepository.DeleteStock(id)
	if err != nil {
		return err
	}
	return nil
}
