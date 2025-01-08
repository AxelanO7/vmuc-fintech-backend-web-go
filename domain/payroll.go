package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Payroll struct {
	ID          uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	Periode     string         `gorm:"not null" json:"periode"`
	Description string         `gorm:"not null" json:"description"`
	Salary      int            `gorm:"not null" json:"salary"`
	Bonus       int            `gorm:"not null" json:"bonus"`
	Penalty     int            `gorm:"not null" json:"penalty"`
	Total       int            `gorm:"not null" json:"total"`
	IdEmployee  uint           `gorm:"not null" json:"id_employee"`
	Employee    *Employee      `json:"employee" gorm:"foreignKey:IdEmployee;references:ID"`
	CreatedAt   *time.Time     `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type PayrollRepository interface {
	RetrieveAllPayroll() ([]Payroll, error)
	RetrievePayrollByID(id uint) (*Payroll, error)
	CreatePayroll(Payroll *Payroll) (*Payroll, error)
	UpdatePayroll(Payroll *Payroll) (*Payroll, error)
	DeletePayroll(id uint) error
}

type PayrollUseCase interface {
	FetchPayrolls(ctx context.Context) ([]Payroll, error)
	FetchPayrollByID(ctx context.Context, id uint) (*Payroll, error)
	AddPayroll(ctx context.Context, req *Payroll) (*Payroll, error)
	EditPayroll(ctx context.Context, req *Payroll) (*Payroll, error)
	DeletePayroll(ctx context.Context, id uint) error
}
