package usecase

import (
	"assyarif-backend-web-go/domain"
	"context"
	"time"
)

type stockOutletUseCase struct {
	stockOutletRepository domain.StockOutletRepository
	contextTimeout        time.Duration
}

func NewStockOutletUseCase(stockOutlet domain.StockOutletRepository, t time.Duration) domain.StockOutletUseCase {
	return &stockOutletUseCase{
		stockOutletRepository: stockOutlet,
		contextTimeout:        t,
	}
}

func (c *stockOutletUseCase) FetchStockOutletByID(ctx context.Context, id uint) (*domain.StockOutlet, error) {
	res, err := c.stockOutletRepository.RetrieveStockOutletByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *stockOutletUseCase) FetchStockOutlets(ctx context.Context) ([]domain.StockOutlet, error) {
	res, err := c.stockOutletRepository.RetrieveAllStockOutlet()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *stockOutletUseCase) CreateStockOutlet(ctx context.Context, req *domain.StockOutlet) (*domain.StockOutlet, error) {
	res, err := c.stockOutletRepository.CreateStockOutlet(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *stockOutletUseCase) UpdateStockOutlet(ctx context.Context, req *domain.StockOutlet) (*domain.StockOutlet, error) {
	res, err := c.stockOutletRepository.UpdateStockOutlet(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *stockOutletUseCase) DeleteStockOutlet(ctx context.Context, id uint) error {
	err := c.stockOutletRepository.DeleteStockOutlet(id)
	if err != nil {
		return err
	}
	return nil
}

func (c *stockOutletUseCase) IncreaseDashboard(ctx context.Context, req *domain.StockOutlet) (*domain.StockOutlet, error) {
	stocks, errStocks := c.stockOutletRepository.RetrieveAllStockOutlet()
	if errStocks != nil {
		return nil, errStocks
	}
	for _, stock := range stocks {
		if stock.IdStuff == req.IdStuff {
			req = &domain.StockOutlet{
				ID:        stock.ID,
				IdStuff:   stock.IdStuff,
				Name:      stock.Name,
				Type:      stock.Type,
				Quantity:  stock.Quantity + req.Quantity,
				Unit:      stock.Unit,
				Price:     stock.Price,
				IdOutlet:  stock.IdOutlet,
				CreatedAt: stock.CreatedAt,
				UpdatedAt: stock.UpdatedAt,
				DeletedAt: stock.DeletedAt,
			}
			res, err := c.stockOutletRepository.UpdateStockOutlet(req)
			if err != nil {
				return nil, err
			}
			return res, nil
		}
	}
	res, err := c.stockOutletRepository.CreateStockOutlet(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *stockOutletUseCase) DecreaseDashboard(ctx context.Context, req *domain.StockOutlet) (*domain.StockOutlet, error) {
	stocks, errStocks := c.stockOutletRepository.RetrieveAllStockOutlet()
	if errStocks != nil {
		return nil, errStocks
	}
	for _, stock := range stocks {
		if stock.IdStuff == req.IdStuff {
			req = &domain.StockOutlet{
				ID:        stock.ID,
				IdStuff:   stock.IdStuff,
				Name:      stock.Name,
				Type:      stock.Type,
				Quantity:  stock.Quantity - req.Quantity,
				Unit:      stock.Unit,
				Price:     stock.Price,
				IdOutlet:  stock.IdOutlet,
				CreatedAt: stock.CreatedAt,
				UpdatedAt: stock.UpdatedAt,
				DeletedAt: stock.DeletedAt,
			}
			res, err := c.stockOutletRepository.UpdateStockOutlet(req)
			if err != nil {
				return nil, err
			}
			return res, nil
		}
	}
	return nil, nil
}

func (c *stockOutletUseCase) IncreaseDashboardMultiple(ctx context.Context, req []domain.StockOutlet) ([]domain.StockOutlet, error) {
	finalStock := []domain.StockOutlet{}
	for _, stock := range req {
		stocks, errStocks := c.stockOutletRepository.RetrieveAllStockOutlet()
		if errStocks != nil {
			return nil, errStocks
		}
		for _, st := range stocks {
			if st.IdStuff == stock.IdStuff {
				stock = domain.StockOutlet{
					ID:        st.ID,
					IdStuff:   st.IdStuff,
					Name:      st.Name,
					Type:      st.Type,
					Quantity:  st.Quantity + stock.Quantity,
					Unit:      st.Unit,
					Price:     st.Price,
					IdOutlet:  st.IdOutlet,
					CreatedAt: st.CreatedAt,
					UpdatedAt: st.UpdatedAt,
					DeletedAt: st.DeletedAt,
				}
				res, err := c.stockOutletRepository.UpdateStockOutlet(&stock)
				if err != nil {
					return nil, err
				}
				_, errCreated := c.stockOutletRepository.CreateStockOutlet(&stock)
				if errCreated != nil {
					return nil, errCreated
				}
				finalStock = append(finalStock, *res)
			}
		}
	}
	return finalStock, nil
}
