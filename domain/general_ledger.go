package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type GeneralLedger struct {
	ID                uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	NameGeneralLedger string         `json:"name_general_ledger"`
	Date              time.Time      `json:"date"`
	CreatedAt         *time.Time     `json:"created_at"`
	UpdatedAt         *time.Time     `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type GeneralLedgerRepository interface {
	RetrieveAllGeneralLedger() ([]GeneralLedger, error)
	RetrieveGeneralLedgerByID(id uint) (*GeneralLedger, error)
	GetGeneralLedgerByGeneralLedgerPeriodeId(id uint) ([]GeneralLedger, error)
	CreateGeneralLedger(req *GeneralLedger) (*GeneralLedger, error)
	CreateBulkGeneralLedger(req []*GeneralLedger) ([]*GeneralLedger, error)
	UpdateGeneralLedger(req *GeneralLedger) (*GeneralLedger, error)
	UpdateBulkGeneralLedger(req []*GeneralLedger) ([]*GeneralLedger, error)
	DeleteGeneralLedger(id uint) error
}

type GeneralLedgerUseCase interface {
	FetchGeneralLedgers(ctx context.Context) ([]GeneralLedger, error)
	FetchGeneralLedgerByID(ctx context.Context, id uint) (*GeneralLedger, error)
	AddGeneralLedger(ctx context.Context, req *GeneralLedger) (*GeneralLedger, error)
	AddBulkGeneralLedger(ctx context.Context, req []*GeneralLedger) ([]*GeneralLedger, error)
	EditGeneralLedger(ctx context.Context, req *GeneralLedger) (*GeneralLedger, error)
	EditBulkGeneralLedger(ctx context.Context, req []*GeneralLedger) ([]*GeneralLedger, error)
	DeleteGeneralLedger(ctx context.Context, id uint) error
}
