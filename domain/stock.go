package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Stock struct {
	ID        uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	Name      string         `json:"name"`
	Type      string         `json:"type"`
	Quantity  float64        `json:"quantity"`
	Unit      string         `json:"unit"`
	Price     float64        `json:"price"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type StockRepository interface {
	RetrieveAllStock() ([]Stock, error)
	RetrieveStockByID(id uint) (*Stock, error)
	CreateStock(Stock *Stock) (*Stock, error)
	UpdateStock(Stock *Stock) (*Stock, error)
	DeleteStock(id uint) error
}

type StockUseCase interface {
	FetchStocks(ctx context.Context) ([]Stock, error)
	FetchStockByID(ctx context.Context, id uint) (*Stock, error)
	CreateStock(ctx context.Context, req *Stock) (*Stock, error)
	UpdateStock(ctx context.Context, req *Stock) (*Stock, error)
	DeleteStock(ctx context.Context, id uint) error
}
