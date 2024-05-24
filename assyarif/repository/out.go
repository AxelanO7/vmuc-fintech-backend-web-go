package repository

import (
	"assyarif-backend-web-go/domain"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type posgreOutRepository struct {
	DB *gorm.DB
}

func NewPostgreOut(client *gorm.DB) domain.OutRepository {
	return &posgreOutRepository{
		DB: client,
	}
}

func (a *posgreOutRepository) RetrieveOuts() ([]domain.Out, error) {
	var res []domain.Out
	err := a.DB.
		Model(domain.Out{}).
		Find(&res).Error
	if err != nil {
		return []domain.Out{}, err
	}
	if len(res) == 0 {
		return []domain.Out{}, errors.New("data not found")
	}
	fmt.Println(res)
	return res, nil
}
