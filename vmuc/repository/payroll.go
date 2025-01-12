package repository

import (
	"errors"
	"fmt"
	"vmuc-fintech-backend-web-go/domain"

	"gorm.io/gorm"
)

type posgrePayrollRepository struct {
	DB *gorm.DB
}

func NewPostgrePayroll(client *gorm.DB) domain.PayrollRepository {
	return &posgrePayrollRepository{
		DB: client,
	}
}

func (a *posgrePayrollRepository) RetrieveAllPayroll() ([]domain.Payroll, error) {
	var res []domain.Payroll
	err := a.DB.
		Model(domain.Payroll{}).
		Preload("User").
		Find(&res).Error
	if err != nil {
		return []domain.Payroll{}, err
	}
	fmt.Println("retrieve all payroll ", res)
	return res, nil
}

func (a *posgrePayrollRepository) RetrievePayrollByID(id uint) (*domain.Payroll, error) {
	var res domain.Payroll
	err := a.DB.
		Model(domain.Payroll{}).
		Where("id = ?", id).
		Preload("User").
		Take(&res).Error
	if err != nil {
		return &domain.Payroll{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.Payroll{}, fmt.Errorf("record not found")
	}
	fmt.Println("retrieve payroll by id ", res)
	return &res, nil
}

func (a *posgrePayrollRepository) CreatePayroll(payroll *domain.Payroll) (*domain.Payroll, error) {
	err := a.DB.
		Model(domain.Payroll{}).
		Create(payroll).Error
	if err != nil {
		return &domain.Payroll{}, err
	}
	fmt.Println("create payroll ", payroll)
	return payroll, nil
}

func (a *posgrePayrollRepository) CreateBulkPayroll(payrolls []*domain.Payroll) ([]*domain.Payroll, error) {
	err := a.DB.
		Model(domain.Payroll{}).
		Create(&payrolls).Error
	if err != nil {
		return []*domain.Payroll{}, err
	}
	fmt.Println("create bulk payroll ", payrolls)
	return payrolls, nil
}

func (a *posgrePayrollRepository) UpdatePayroll(payroll *domain.Payroll) (*domain.Payroll, error) {
	err := a.DB.
		Model(domain.Payroll{}).
		Where("id = ?", payroll.ID).
		Updates(payroll).Error
	if err != nil {
		return &domain.Payroll{}, err
	}
	fmt.Println("update payroll ", payroll)
	return payroll, nil
}

func (a *posgrePayrollRepository) UpdateBulkPayroll(payrolls []*domain.Payroll) ([]*domain.Payroll, error) {
	err := a.DB.
		Model(domain.Payroll{}).
		Updates(&payrolls).Error
	if err != nil {
		return []*domain.Payroll{}, err
	}
	fmt.Println("update bulk payroll ", payrolls)
	return payrolls, nil
}

func (a *posgrePayrollRepository) DeletePayroll(id uint) error {
	err := a.DB.
		Model(domain.Payroll{}).
		Where("id = ?", id).
		Delete(&domain.Payroll{}).Error
	if err != nil {
		return err
	}
	fmt.Println("delete payroll ", id)
	return nil
}
