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

func (a *posgreInRepository) RetrieveInLastNumber() (int, error) {
	var res []domain.In
	a.DB.
		Model(domain.In{}).
		Find(&res)

	lastNumber := 0
	for _, v := range res {
		fmt.Println(v.ID)
		if v.ID > uint(lastNumber) {
			lastNumber = int(v.ID)
		}

	}

	fmt.Println(lastNumber)
	return lastNumber, nil
}

func (a *posgreInRepository) CreateIn(in domain.In) (domain.In, error) {
	err := a.DB.
		Model(domain.In{}).
		Create(&in).Error
	if err != nil {
		return domain.In{}, err
	}
	fmt.Println(in)
	return in, nil
}

func (a *posgreInRepository) RetrieveInById(id string) (domain.In, error) {
	var res domain.In
	err := a.DB.
		Model(domain.In{}).
		Where("id = ?", id).
		First(&res).Error
	if err != nil {
		return domain.In{}, err
	}
	fmt.Println(res)
	return res, nil
}

func (a *posgreInRepository) UpdateInById(in domain.In) (domain.In, error) {
	err := a.DB.
		Model(domain.In{}).
		Where("id = ?", in.ID).
		Updates(&in).Error
	if err != nil {
		return domain.In{}, err
	}
	fmt.Println(in)
	return in, nil
}

func (a *posgreInRepository) RemoveInById(id string) error {
	err := a.DB.
		Model(domain.In{}).
		Where("id = ?", id).
		Delete(&domain.In{}).Error
	if err != nil {
		return err
	}
	return nil
}
