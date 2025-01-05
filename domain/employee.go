package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	ID        uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	Address   string         `gorm:"not null" json:"address"`
	Phone     string         `gorm:"not null" json:"phone"`
	Position  string         `gorm:"not null" json:"position"`
	IdUser    uint           `gorm:"not null" json:"id_user"`
	User      *User          `json:"user" gorm:"foreignKey:IdUser;references:ID"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type EmployeeRepository interface {
	RetrieveAllEmployee() ([]Employee, error)
	RetrieveEmployeeByID(id uint) (*Employee, error)
	CreateEmployee(Employee *Employee) (*Employee, error)
	UpdateEmployee(Employee *Employee) (*Employee, error)
	DeleteEmployee(id uint) error
}

type EmployeeUseCase interface {
	FetchEmployees(ctx context.Context) ([]Employee, error)
	FetchEmployeeByID(ctx context.Context, id uint) (*Employee, error)
	CreateEmployee(ctx context.Context, req *Employee) (*Employee, error)
	UpdateEmployee(ctx context.Context, req *Employee) (*Employee, error)
	DeleteEmployee(ctx context.Context, id uint) error
	ShowEmployeeLastNumber(ctx context.Context) (int, error)
}
