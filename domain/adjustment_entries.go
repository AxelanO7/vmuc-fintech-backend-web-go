package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type AdjusmentEntries struct {
	ID             uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	NameAccount    string         `json:"name_account"`
	Date           time.Time      `json:"date"`
	IdRef          int            `gorm:"not null" json:"id_ref"`
	Information    string         `json:"information"`
	Debit          float64        `json:"debit"`
	Kredit         float64        `json:"kredit"`
	IdPeriode      uint           `gorm:"not null" json:"id_periode"`
	PayrollPeriode *Periode       `json:"payroll_periode" gorm:"foreignKey:IdPayrollPeriode;references:ID"`
	Ref            *Ref           `json:"ref" gorm:"foreignKey:IdRef;references:ID"`
	CreatedAt      *time.Time     `json:"created_at"`
	UpdatedAt      *time.Time     `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type AdjusmentEntriesRepository interface {
	RetrieveAllAdjusmentEntries() ([]AdjusmentEntries, error)
	RetrieveAdjusmentEntriesByID(id uint) (*AdjusmentEntries, error)
	GetAdjusmentEntriesByAdjusmentEntriesPeriodeId(id uint) ([]AdjusmentEntries, error)
	CreateAdjusmentEntries(req *AdjusmentEntries) (*AdjusmentEntries, error)
	CreateBulkAdjusmentEntries(req []*AdjusmentEntries) ([]*AdjusmentEntries, error)
	UpdateAdjusmentEntries(req *AdjusmentEntries) (*AdjusmentEntries, error)
	UpdateBulkAdjusmentEntries(req []*AdjusmentEntries) ([]*AdjusmentEntries, error)
	DeleteAdjusmentEntries(id uint) error
}

type AdjusmentEntriesUseCase interface {
	FetchAdjusmentEntriess(ctx context.Context) ([]AdjusmentEntries, error)
	FetchAdjusmentEntriesByID(ctx context.Context, id uint) (*AdjusmentEntries, error)
	AddAdjusmentEntries(ctx context.Context, req *AdjusmentEntries) (*AdjusmentEntries, error)
	AddBulkAdjusmentEntries(ctx context.Context, req []*AdjusmentEntries) ([]*AdjusmentEntries, error)
	EditAdjusmentEntries(ctx context.Context, req *AdjusmentEntries) (*AdjusmentEntries, error)
	EditBulkAdjusmentEntries(ctx context.Context, req []*AdjusmentEntries) ([]*AdjusmentEntries, error)
	DeleteAdjusmentEntries(ctx context.Context, id uint) error
}
