package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Out struct {
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

type OutRepository interface {
	RetrieveOuts() ([]Out, error)
	RetrieveOutById(id string) (Out, error)
	RetrieveOutLastNumber() (int, error)
	CreateOut(out Out) (Out, error)
	UpdateOutById(out Out) (Out, error)
	RemoveOutById(id string) error
}

type OutUseCase interface {
	GetOuts(ctx context.Context) ([]Out, error)
	ShowOutById(ctx context.Context, id string) (Out, error)
	ShowOutLastNumber(ctx context.Context) (int, error)
	AddOut(ctx context.Context, out Out) (Out, error)
	EditOutById(ctx context.Context, out Out) (Out, error)
	DeleteOutById(ctx context.Context, id string) error
}
