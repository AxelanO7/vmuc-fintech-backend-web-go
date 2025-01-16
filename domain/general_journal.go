package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type GeneralJournal struct {
	ID             uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	NameAccount    string         `json:"name_account"`
	Date           time.Time      `json:"date"`
	IdRef          int            `gorm:"not null" json:"id_ref"`
	Information    string         `json:"information"`
	Debit          float64        `json:"debit"`
	Kredit         float64        `json:"kredit"`
	Ref            *Ref           `json:"ref" gorm:"foreignKey:IdRef;references:ID"`
	IdPeriode      uint           `gorm:"not null" json:"id_periode"`
	PayrollPeriode *Periode       `json:"payroll_periode" gorm:"foreignKey:IdPayrollPeriode;references:ID"`
	CreatedAt      *time.Time     `json:"created_at"`
	UpdatedAt      *time.Time     `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type GeneralJournalRepository interface {
	RetrieveAllGeneralJournal() ([]GeneralJournal, error)
	RetrieveGeneralJournalByID(id uint) (*GeneralJournal, error)
	GetGeneralJournalByGeneralJournalPeriodeId(id uint) ([]GeneralJournal, error)
	CreateGeneralJournal(req *GeneralJournal) (*GeneralJournal, error)
	CreateBulkGeneralJournal(req []*GeneralJournal) ([]*GeneralJournal, error)
	UpdateGeneralJournal(req *GeneralJournal) (*GeneralJournal, error)
	UpdateBulkGeneralJournal(req []*GeneralJournal) ([]*GeneralJournal, error)
	DeleteGeneralJournal(id uint) error
}

type GeneralJournalUseCase interface {
	FetchGeneralJournals(ctx context.Context) ([]GeneralJournal, error)
	FetchGeneralJournalByID(ctx context.Context, id uint) (*GeneralJournal, error)
	AddGeneralJournal(ctx context.Context, req *GeneralJournal) (*GeneralJournal, error)
	AddBulkGeneralJournal(ctx context.Context, req []*GeneralJournal) ([]*GeneralJournal, error)
	EditGeneralJournal(ctx context.Context, req *GeneralJournal) (*GeneralJournal, error)
	EditBulkGeneralJournal(ctx context.Context, req []*GeneralJournal) ([]*GeneralJournal, error)
	DeleteGeneralJournal(ctx context.Context, id uint) error
}
