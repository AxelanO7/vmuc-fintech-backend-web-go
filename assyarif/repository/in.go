package repository

import (
	"assyarif-backend-web-go/domain"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type posgreInRepository struct {
	DB *gorm.DB
}

func NewPostgreIn(client *gorm.DB) domain.InRepository {
	return &posgreInRepository{
		DB: client,
	}
}

func (a *posgreInRepository) RetrieveIns() ([]domain.In, error) {
	var res []domain.In
	err := a.DB.
		Model(domain.In{}).
		Find(&res).Error
	if err != nil {
		return []domain.In{}, err
	}
	if len(res) == 0 {
		return []domain.In{}, errors.New("data not found")
	}
	fmt.Println(res)
	return res, nil
}
