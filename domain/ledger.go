package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Ledger struct {
	ID                uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	NameGeneralLedger string         `json:"name_general_ledger"`
	Date              string         `json:"date"`
	CreatedAt         *time.Time     `json:"created_at"`
	UpdatedAt         *time.Time     `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type LedgerRepository interface {
	RetrieveLedgers() ([]Ledger, error)
	RetrieveLedgerByID(id uint) (*Ledger, error)
	GetLedgerByPeriodeId(id uint) ([]Ledger, error)
	CreateLedger(req *Ledger) (*Ledger, error)
	CreateBulkLedger(req []*Ledger) ([]*Ledger, error)
	UpdateLedger(req *Ledger) (*Ledger, error)
	UpdateBulkLedger(req []*Ledger) ([]*Ledger, error)
	DeleteLedger(id uint) error
}

type LedgerUseCase interface {
	FetchLedgers(ctx context.Context) ([]Ledger, error)
	FetchLedgerByID(ctx context.Context, id uint, opt bool) (map[string]any, error)
	AddLedger(ctx context.Context, req *Ledger) (*Ledger, error)
	AddBulkLedger(ctx context.Context, req []*Ledger) ([]*Ledger, error)
	EditLedger(ctx context.Context, req *Ledger) (*Ledger, error)
	EditBulkLedger(ctx context.Context, req []*Ledger) ([]*Ledger, error)
	DeleteLedger(ctx context.Context, id uint) error
}
