package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Outlet struct {
	ID        uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	Address   string         `gorm:"not null" json:"address"`
	Phone     string         `gorm:"not null" json:"phone"`
	IdUser    uint           `gorm:"not null" json:"id_user"`
	User      *User          `json:"user" gorm:"foreignKey:IdUser;references:ID"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type OutletRepository interface {
	RetrieveAllOutlet() ([]Outlet, error)
	RetrieveOutletByID(id uint) (*Outlet, error)
	CreateOutlet(Outlet *Outlet) (*Outlet, error)
	UpdateOutlet(Outlet *Outlet) (*Outlet, error)
	DeleteOutlet(id uint) error
	ShowOutletByIDUser(id uint) (*Outlet, error)
}

type OutletUseCase interface {
	FetchOutlets(ctx context.Context) ([]Outlet, error)
	FetchOutletByID(ctx context.Context, id uint) (*Outlet, error)
	CreateOutlet(ctx context.Context, req *Outlet) (*Outlet, error)
	UpdateOutlet(ctx context.Context, req *Outlet) (*Outlet, error)
	DeleteOutlet(ctx context.Context, id uint) error
	ShowOutletLastNumber(ctx context.Context) (int, error)
	ShowOutletByIDUser(ctx context.Context, id uint) (*Outlet, error)
}
