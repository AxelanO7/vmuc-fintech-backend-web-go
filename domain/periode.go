package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Periode struct {
	ID               uint               `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	Period           string             `gorm:"not null" json:"period"`
	Description      string             `gorm:"not null" json:"description"`
	Payrolls         []Payroll          `gorm:"-" json:"payroll"`
	AdjusmentEntries []AdjusmentEntries `gorm:"-" json:"adjusment_entries"`
	GeneralJournal   []GeneralJournal   `gorm:"-" json:"general_journal"`
	TrialBalance     []TrialBalance     `gorm:"-" json:"trial_balance"`
	CreatedAt        *time.Time         `json:"created_at"`
	UpdatedAt        *time.Time         `json:"updated_at"`
	DeletedAt        gorm.DeletedAt     `gorm:"index" json:"deleted_at"`
}

type PeriodeRepository interface {
	RetrieveAllPeriode() ([]Periode, error)
	RetrievePeriodeByID(id uint) (*Periode, error)
	CreatePeriode(Periode *Periode) (*Periode, error)
	CreateBulkPeriode(Periode []*Periode) ([]*Periode, error)
	UpdatePeriode(Periode *Periode) (*Periode, error)
	UpdateBulkPeriode(Periode []*Periode) ([]*Periode, error)
	DeletePeriode(id uint) error
}

type PeriodeUseCase interface {
	FetchPayrollPeriode(ctx context.Context) ([]Periode, error)
	FetchAdjusmentEntriesPeriode(ctx context.Context) ([]Periode, error)
	FetchGeneralJournalPeriode(ctx context.Context) ([]Periode, error)
	FetchTrialBalancePeriode(ctx context.Context) ([]Periode, error)
	FetchPeriodeByID(ctx context.Context, id uint) (*Periode, error)
	AddPeriode(ctx context.Context, req *Periode) (*Periode, error)
	AddBulkPeriode(ctx context.Context, req []*Periode) ([]*Periode, error)
	EditPeriode(ctx context.Context, req *Periode) (*Periode, error)
	EditBulkPeriode(ctx context.Context, req []*Periode) ([]*Periode, error)
	DeletePeriode(ctx context.Context, id uint) error
	AddPayrollPeriode(ctx context.Context, req *PayrollPeriode) (*PayrollPeriode, error)
}
