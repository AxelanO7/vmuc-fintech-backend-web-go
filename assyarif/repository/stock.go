package repository

import (
	"assyarif-backend-web-go/domain"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type posgreStockRepository struct {
	DB *gorm.DB
}

func NewPostgreStock(client *gorm.DB) domain.StockRepository {
	return &posgreStockRepository{
		DB: client,
	}
}

func (a *posgreStockRepository) RetrieveAllStock() ([]domain.Stock, error) {
	var res []domain.Stock
	err := a.DB.
		Model(domain.Stock{}).
		Find(&res).Error
	if err != nil {
		return []domain.Stock{}, err
	}
	fmt.Println(res)
	return res, nil
}

func (a *posgreStockRepository) RetrieveStockByID(id uint) (*domain.Stock, error) {
	var res domain.Stock
	err := a.DB.
		Model(domain.Stock{}).
		Where("id = ?", id).
		Take(&res).Error
	if err != nil {
		return &domain.Stock{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.Stock{}, fmt.Errorf("record not found")
	}
	fmt.Println(res)
	return &res, nil
}

func (a *posgreStockRepository) CreateStock(user *domain.Stock) (*domain.Stock, error) {
	err := a.DB.
		Model(domain.Stock{}).
		Create(user).Error
	if err != nil {
		return &domain.Stock{}, err
	}
	fmt.Println(user)
	return user, nil
}

func (a *posgreStockRepository) UpdateStock(user *domain.Stock) (*domain.Stock, error) {
	err := a.DB.
		Model(domain.Stock{}).
		Where("id = ?", user.ID).
		Updates(user).Error
	if err != nil {
		return &domain.Stock{}, err
	}
	fmt.Println(user)
	return user, nil
}

func (a *posgreStockRepository) DeleteStock(id uint) error {
	err := a.DB.
		Model(domain.Stock{}).
		Where("id = ?", id).
		Delete(&domain.Stock{}).Error
	if err != nil {
		return err
	}
	return nil
}
