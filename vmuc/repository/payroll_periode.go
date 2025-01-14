package repository

import (
	"errors"
	"fmt"
	"vmuc-fintech-backend-web-go/domain"

	"gorm.io/gorm"
)

type posgrePayrollPeriodeRepository struct {
	DB *gorm.DB
}

func NewPostgrePayrollPeriode(client *gorm.DB) domain.PayrollPeriodeRepository {
	return &posgrePayrollPeriodeRepository{
		DB: client,
	}
}

func (a *posgrePayrollPeriodeRepository) RetrieveAllPayrollPeriode() ([]domain.PayrollPeriode, error) {
	var res []domain.PayrollPeriode
	err := a.DB.
		Model(domain.PayrollPeriode{}).
		Find(&res).Error
	if err != nil {
		return []domain.PayrollPeriode{}, err
	}
	fmt.Println("retrieve all payroll ", res)
	return res, nil
}

func (a *posgrePayrollPeriodeRepository) RetrievePayrollPeriodeByID(id uint) (*domain.PayrollPeriode, error) {
	var res domain.PayrollPeriode
	err := a.DB.
		Model(domain.PayrollPeriode{}).
		Where("id = ?", id).
		Preload("Payroll").
		Take(&res).Error
	if err != nil {
		return &domain.PayrollPeriode{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.PayrollPeriode{}, fmt.Errorf("record not found")
	}
	fmt.Println("retrieve payroll by id ", res)
	return &res, nil
}

func (a *posgrePayrollPeriodeRepository) CreatePayrollPeriode(payrollPeriode *domain.PayrollPeriode) (*domain.PayrollPeriode, error) {
	err := a.DB.
		Model(domain.PayrollPeriode{}).
		Create(payrollPeriode).Error
	if err != nil {
		return &domain.PayrollPeriode{}, err
	}
	fmt.Println("create payroll ", payrollPeriode)
	return payrollPeriode, nil
}

func (a *posgrePayrollPeriodeRepository) CreateBulkPayrollPeriode(payrolls []*domain.PayrollPeriode) ([]*domain.PayrollPeriode, error) {
	err := a.DB.
		Model(domain.PayrollPeriode{}).
		Create(&payrolls).Error
	if err != nil {
		return []*domain.PayrollPeriode{}, err
	}
	fmt.Println("create bulk payroll ", payrolls)
	return payrolls, nil
}

func (a *posgrePayrollPeriodeRepository) UpdatePayrollPeriode(payroll *domain.PayrollPeriode) (*domain.PayrollPeriode, error) {
	err := a.DB.
		Model(domain.PayrollPeriode{}).
		Where("id = ?", payroll.ID).
		Updates(payroll).Error
	if err != nil {
		return &domain.PayrollPeriode{}, err
	}
	fmt.Println("update payroll ", payroll)
	return payroll, nil
}

func (a *posgrePayrollPeriodeRepository) UpdateBulkPayrollPeriode(payrolls []*domain.PayrollPeriode) ([]*domain.PayrollPeriode, error) {
	err := a.DB.
		Model(domain.PayrollPeriode{}).
		Updates(&payrolls).Error
	if err != nil {
		return []*domain.PayrollPeriode{}, err
	}
	fmt.Println("update bulk payroll ", payrolls)
	return payrolls, nil
}

func (a *posgrePayrollPeriodeRepository) DeletePayrollPeriode(id uint) error {
	err := a.DB.
		Model(domain.PayrollPeriode{}).
		Where("id = ?", id).
		Delete(&domain.PayrollPeriode{}).Error
	if err != nil {
		return err
	}
	fmt.Println("delete payroll ", id)
	return nil
}
