package repository

import (
	"errors"
	"fmt"
	"vmuc-fintech-backend-web-go/domain"

	"gorm.io/gorm"
)

type posgreEmployeeRepository struct {
	DB *gorm.DB
}

func NewPostgreEmployee(client *gorm.DB) domain.EmployeeRepository {
	return &posgreEmployeeRepository{
		DB: client,
	}
}

func (a *posgreEmployeeRepository) RetrieveAllEmployee() ([]domain.Employee, error) {
	var res []domain.Employee
	err := a.DB.
		Model(domain.Employee{}).
		Find(&res).Error
	if err != nil {
		return []domain.Employee{}, err
	}
	fmt.Println("retrieve all employee ", res)
	return res, nil
}

func (a *posgreEmployeeRepository) RetrieveEmployeeByID(id uint) (*domain.Employee, error) {
	var res domain.Employee
	err := a.DB.
		Model(domain.Employee{}).
		Where("id = ?", id).
		Preload("User").
		Take(&res).Error
	if err != nil {
		return &domain.Employee{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.Employee{}, fmt.Errorf("record not found")
	}
	fmt.Println("retrieve employee by id ", res)
	return &res, nil
}

func (a *posgreEmployeeRepository) CreateEmployee(employee *domain.Employee) (*domain.Employee, error) {
	err := a.DB.
		Model(domain.Employee{}).
		Create(employee).Error
	if err != nil {
		return &domain.Employee{}, err
	}
	fmt.Println("create employee ", employee)
	return employee, nil
}

func (a *posgreEmployeeRepository) UpdateEmployee(employee *domain.Employee) (*domain.Employee, error) {
	err := a.DB.
		Model(domain.Employee{}).
		Where("id = ?", employee.ID).
		Updates(employee).Error
	if err != nil {
		return &domain.Employee{}, err
	}
	fmt.Println("update employee ", employee)
	return employee, nil
}

func (a *posgreEmployeeRepository) DeleteEmployee(id uint) error {
	err := a.DB.
		Model(domain.Employee{}).
		Where("id = ?", id).
		Delete(&domain.Employee{}).Error
	if err != nil {
		return err
	}
	fmt.Println("delete employee ", id)
	return nil
}
