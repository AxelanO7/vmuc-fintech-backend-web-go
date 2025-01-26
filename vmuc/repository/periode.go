package repository

import (
	"errors"
	"fmt"
	"vmuc-fintech-backend-web-go/domain"

	"gorm.io/gorm"
)

type posgrePeriodeRepository struct {
	DB *gorm.DB
}

func NewPostgrePeriode(client *gorm.DB) domain.PeriodeRepository {
	return &posgrePeriodeRepository{
		DB: client,
	}
}

func (a *posgrePeriodeRepository) RetrieveAllPeriode() ([]domain.Periode, error) {
	var res []domain.Periode
	err := a.DB.
		Model(domain.Periode{}).
		Find(&res).Error
	if err != nil {
		return []domain.Periode{}, err
	}
	fmt.Println("retrieve all payroll ", res)
	return res, nil
}

func (a *posgrePeriodeRepository) GetPeriodeByPeriode(periode string) (*domain.Periode, error) {
	var res domain.Periode
	err := a.DB.
		Model(domain.Periode{}).
		Where("period = ?", periode).
		Take(&res).Error
	if err != nil {
		return &domain.Periode{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.Periode{}, fmt.Errorf("record not found")
	}
	fmt.Println("retrieve payroll by id ", res)
	return &res, nil
}

func (a *posgrePeriodeRepository) RetrievePeriodeByID(id uint) (*domain.Periode, error) {
	var res domain.Periode
	err := a.DB.
		Model(domain.Periode{}).
		Where("id = ?", id).
		Preload("Payroll").
		Take(&res).Error
	if err != nil {
		return &domain.Periode{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.Periode{}, fmt.Errorf("record not found")
	}
	fmt.Println("retrieve payroll by id ", res)
	return &res, nil
}

func (a *posgrePeriodeRepository) CreatePeriode(Periode *domain.Periode) (*domain.Periode, error) {
	err := a.DB.
		Model(domain.Periode{}).
		Create(Periode).Error
	if err != nil {
		return &domain.Periode{}, err
	}
	fmt.Println("create payroll ", Periode)
	return Periode, nil
}

func (a *posgrePeriodeRepository) CreateBulkPeriode(payrolls []*domain.Periode) ([]*domain.Periode, error) {
	err := a.DB.
		Model(domain.Periode{}).
		Create(&payrolls).Error
	if err != nil {
		return []*domain.Periode{}, err
	}
	fmt.Println("create bulk payroll ", payrolls)
	return payrolls, nil
}

func (a *posgrePeriodeRepository) UpdatePeriode(payroll *domain.Periode) (*domain.Periode, error) {
	err := a.DB.
		Model(domain.Periode{}).
		Where("id = ?", payroll.ID).
		Updates(payroll).Error
	if err != nil {
		return &domain.Periode{}, err
	}
	fmt.Println("update payroll ", payroll)
	return payroll, nil
}

func (a *posgrePeriodeRepository) UpdateBulkPeriode(payrolls []*domain.Periode) ([]*domain.Periode, error) {
	err := a.DB.
		Model(domain.Periode{}).
		Updates(&payrolls).Error
	if err != nil {
		return []*domain.Periode{}, err
	}
	fmt.Println("update bulk payroll ", payrolls)
	return payrolls, nil
}

func (a *posgrePeriodeRepository) DeletePeriode(id uint) error {
	err := a.DB.
		Model(domain.Periode{}).
		Where("id = ?", id).
		Delete(&domain.Periode{}).Error
	if err != nil {
		return err
	}
	fmt.Println("delete payroll ", id)
	return nil
}
