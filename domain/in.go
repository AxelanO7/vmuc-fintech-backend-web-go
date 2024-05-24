package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type In struct {
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

type InRepository interface {
	RetrieveIns() ([]In, error)
}

type InUseCase interface {
	GetIns(ctx context.Context) ([]In, error)
}
