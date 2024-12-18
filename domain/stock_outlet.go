package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type StockOutlet struct {
	ID        uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	IdStuff   uint           `json:"id_stuff"`
	IdOut     uint           `json:"id_out"`
	Out       *Out           `gorm:"foreignKey:IdOut;references:ID" json:"out"`
	IdOutlet  uint           `json:"id_outlet"`
	Outlet    *Outlet        `gorm:"foreignKey:IdOutlet;references:ID" json:"outlet"`
	Name      string         `json:"name"`
	Type      string         `json:"type"`
	Quantity  float64        `json:"quantity"`
	Unit      string         `json:"unit"`
	Price     float64        `json:"price"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type StockOutlets struct {
	StockOutlets []StockOutlet `json:"orders"`
}

type StockOutletRepository interface {
	RetrieveAllStockOutlet() ([]StockOutlet, error)
	RetrieveStockOutletByID(id uint) (*StockOutlet, error)
	CreateStockOutlet(StockOutlet *StockOutlet) (*StockOutlet, error)
	UpdateStockOutlet(StockOutlet *StockOutlet) (*StockOutlet, error)
	DeleteStockOutlet(id uint) error
	UpdateStockOutletsMultiple(StockOutlets []StockOutlet) ([]StockOutlet, error)
	CreateStockOutletsMultiple(StockOutlets []StockOutlet) ([]StockOutlet, error)
}

type StockOutletUseCase interface {
	FetchStockOutlets(ctx context.Context) ([]StockOutlet, error)
	FetchStockOutletByID(ctx context.Context, id uint) (*StockOutlet, error)
	CreateStockOutlet(ctx context.Context, req *StockOutlet) (*StockOutlet, error)
	UpdateStockOutlet(ctx context.Context, req *StockOutlet) (*StockOutlet, error)
	DeleteStockOutlet(ctx context.Context, id uint) error
	IncreaseDashboard(ctx context.Context, req *StockOutlet) (*StockOutlet, error)
	DecreaseDashboard(ctx context.Context, req *StockOutlet) (*StockOutlet, error)
	IncreaseDashboardMultiple(ctx context.Context, req []StockOutlet) ([]StockOutlet, error)
}
