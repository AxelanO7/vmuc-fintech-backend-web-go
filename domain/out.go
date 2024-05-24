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
}

type OutUseCase interface {
	GetOuts(ctx context.Context) ([]Out, error)
}
