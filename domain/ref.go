package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Ref struct {
	ID        uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	Type      string         `gorm:"not null" json:"type"`
	Code      uint           `gorm:"not null" json:"code"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type RefRepository interface {
	RetrieveRefs() ([]Ref, error)
	RetrieveRefByID(id uint) (*Ref, error)
	CreateRef(Ref *Ref) (*Ref, error)
	UpdateRef(Ref *Ref) (*Ref, error)
	DeleteRef(id uint) error
}

type RefUseCase interface {
	FetchRefs(ctx context.Context) ([]Ref, error)
	FetchRefByID(ctx context.Context, id uint) (*Ref, error)
	AddRef(ctx context.Context, req *Ref) (*Ref, error)
	EditRef(ctx context.Context, req *Ref) (*Ref, error)
	DeleteRef(ctx context.Context, id uint) error
}
