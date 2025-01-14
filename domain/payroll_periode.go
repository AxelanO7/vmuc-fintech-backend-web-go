package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type PayrollPeriode struct {
	ID          uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	Period      string         `gorm:"not null" json:"period"`
	Description string         `gorm:"not null" json:"description"`
	Payrolls    []Payroll      `gorm:"-" json:"payroll"`
	CreatedAt   *time.Time     `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type PayrollPeriodeRepository interface {
	RetrieveAllPayrollPeriode() ([]PayrollPeriode, error)
	RetrievePayrollPeriodeByID(id uint) (*PayrollPeriode, error)
	CreatePayrollPeriode(payrollPeriode *PayrollPeriode) (*PayrollPeriode, error)
	CreateBulkPayrollPeriode(payrollPeriode []*PayrollPeriode) ([]*PayrollPeriode, error)
	UpdatePayrollPeriode(payrollPeriode *PayrollPeriode) (*PayrollPeriode, error)
	UpdateBulkPayrollPeriode(payrollPeriode []*PayrollPeriode) ([]*PayrollPeriode, error)
	DeletePayrollPeriode(id uint) error
}

type PayrollPeriodeUseCase interface {
	FetchPayrollPeriode(ctx context.Context) ([]PayrollPeriode, error)
	FetchPayrollPeriodeByID(ctx context.Context, id uint) (*PayrollPeriode, error)
	AddPayrollPeriode(ctx context.Context, req *PayrollPeriode) (*PayrollPeriode, error)
	AddBulkPayrollPeriode(ctx context.Context, req []*PayrollPeriode) ([]*PayrollPeriode, error)
	EditPayrollPeriode(ctx context.Context, req *PayrollPeriode) (*PayrollPeriode, error)
	EditBulkPayrollPeriode(ctx context.Context, req []*PayrollPeriode) ([]*PayrollPeriode, error)
	DeletePayrollPeriode(ctx context.Context, id uint) error
}
