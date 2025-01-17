package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type TrialBalance struct {
	ID          uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	NameAccount string         `json:"name_account"`
	IdRef       int            `gorm:"not null" json:"id_ref"`
	Debit       float64        `json:"debit"`
	Kredit      float64        `json:"kredit"`
	IdPeriode   uint           `gorm:"not null" json:"id_periode"`
	Periode     *Periode       `json:"payroll_periode" gorm:"foreignKey:IdPeriode;references:ID"`
	Ref         *Ref           `json:"ref" gorm:"foreignKey:IdRef;references:ID"`
	CreatedAt   *time.Time     `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type TrialBalanceRepository interface {
	RetrieveAllTrialBalance() ([]TrialBalance, error)
	RetrieveTrialBalanceByID(id uint) (*TrialBalance, error)
	GetTrialBalanceByTrialBalancePeriodeId(id uint) ([]TrialBalance, error)
	CreateTrialBalance(req *TrialBalance) (*TrialBalance, error)
	CreateBulkTrialBalance(req []*TrialBalance) ([]*TrialBalance, error)
	UpdateTrialBalance(req *TrialBalance) (*TrialBalance, error)
	UpdateBulkTrialBalance(req []*TrialBalance) ([]*TrialBalance, error)
	DeleteTrialBalance(id uint) error
}

type TrialBalanceUseCase interface {
	FetchTrialBalances(ctx context.Context) ([]TrialBalance, error)
	FetchTrialBalanceByID(ctx context.Context, id uint) (*TrialBalance, error)
	AddTrialBalance(ctx context.Context, req *TrialBalance) (*TrialBalance, error)
	AddBulkTrialBalance(ctx context.Context, req []*TrialBalance) ([]*TrialBalance, error)
	EditTrialBalance(ctx context.Context, req *TrialBalance) (*TrialBalance, error)
	EditBulkTrialBalance(ctx context.Context, req []*TrialBalance) ([]*TrialBalance, error)
	DeleteTrialBalance(ctx context.Context, id uint) error
}
