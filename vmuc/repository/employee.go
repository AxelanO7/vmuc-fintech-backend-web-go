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
		Preload("User").
		Find(&res).Error
	if err != nil {
		return []domain.Employee{}, err
	}
	fmt.Println(res)
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
	fmt.Println(res)
	return &res, nil
}

func (a *posgreEmployeeRepository) CreateEmployee(user *domain.Employee) (*domain.Employee, error) {
	err := a.DB.
		Model(domain.Employee{}).
		Create(user).Error
	if err != nil {
		return &domain.Employee{}, err
	}
	fmt.Println(user)
	return user, nil
}

func (a *posgreEmployeeRepository) UpdateEmployee(user *domain.Employee) (*domain.Employee, error) {
	err := a.DB.
		Model(domain.Employee{}).
		Where("id = ?", user.ID).
		Updates(user).Error
	if err != nil {
		return &domain.Employee{}, err
	}
	fmt.Println(user)
	return user, nil
}

func (a *posgreEmployeeRepository) DeleteEmployee(id uint) error {
	err := a.DB.
		Model(domain.Employee{}).
		Where("id = ?", id).
		Delete(&domain.Employee{}).Error
	if err != nil {
		return err
	}
	return nil
}
