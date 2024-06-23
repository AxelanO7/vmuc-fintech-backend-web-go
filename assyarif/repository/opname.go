package repository

import (
	"assyarif-backend-web-go/domain"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type posgreOpnameRepository struct {
	DB *gorm.DB
}

func NewPostgreOpname(client *gorm.DB) domain.OpnameRepository {
	return &posgreOpnameRepository{
		DB: client,
	}
}

func (a *posgreOpnameRepository) CreateOpname(opname *domain.Opname) error {
	err := a.DB.
		Model(domain.Opname{}).
		Create(opname).Error
	if err != nil {
		return err
	}
	fmt.Println(opname)
	return nil
}

func (a *posgreOpnameRepository) RetrieveAllOpname() ([]domain.Opname, error) {
	var res []domain.Opname
	err := a.DB.
		Model(domain.Opname{}).
		Find(&res).Error
	if err != nil {
		return []domain.Opname{}, err
	}
	fmt.Println(res)
	return res, nil
}

func (a *posgreOpnameRepository) RetrieveOpnameByID(id uint) (*domain.Opname, error) {
	var res domain.Opname
	err := a.DB.
		Model(domain.Opname{}).
		Where("id = ?", id).
		Take(&res).Error
	if err != nil {
		return &domain.Opname{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.Opname{}, fmt.Errorf("record not found")
	}
	fmt.Println(res)
	return &res, nil
}

func (a *posgreOpnameRepository) RetriveByStartDateEndDate(startDate, endDate string) ([]domain.In, []domain.Out, []domain.Rtr, error) {
	var resIn []domain.In
	var resOut []domain.Out
	var resRtr []domain.Rtr
	err := a.DB.
		Model(domain.In{}).
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Find(&resIn).Error
	if err != nil {
		return []domain.In{}, []domain.Out{}, []domain.Rtr{}, err
	}
	err = a.DB.
		Model(domain.Out{}).
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Preload("Order").
		Preload("Order.Stock").
		Find(&resOut).Error
	if err != nil {
		return []domain.In{}, []domain.Out{}, []domain.Rtr{}, err
	}
	err = a.DB.
		Model(domain.Rtr{}).
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Preload("Stock").
		Find(&resRtr).Error
	if err != nil {
		return []domain.In{}, []domain.Out{}, []domain.Rtr{}, err
	}
	return resIn, resOut, resRtr, nil
}
