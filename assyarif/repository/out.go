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

func (a *posgreOutRepository) RetrieveOutLastNumber() (int, error) {
	var res []domain.Out
	a.DB.
		Model(domain.Out{}).
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

func (a *posgreOutRepository) CreateOut(out domain.Out) (domain.Out, error) {
	err := a.DB.
		Model(domain.Out{}).
		Create(&out).Error
	if err != nil {
		return domain.Out{}, err
	}
	return out, nil
}

func (a *posgreOutRepository) RetrieveOutByID(id string) (domain.Out, error) {
	var res domain.Out
	err := a.DB.
		Model(domain.Out{}).
		Where("id = ?", id).
		First(&res).Error
	if err != nil {
		return domain.Out{}, err
	}
	return res, nil
}

func (a *posgreOutRepository) UpdateOutByID(out domain.Out) (domain.Out, error) {
	err := a.DB.
		Model(domain.Out{}).
		Where("id = ?", out.ID).
		Updates(&out).Error
	if err != nil {
		return domain.Out{}, err
	}
	return out, nil
}

func (a *posgreOutRepository) RemoveOutByID(id string) error {
	var res domain.Out
	err := a.DB.
		Model(domain.Out{}).
		Where("id = ?", id).
		Delete(&res).Error
	if err != nil {
		return err
	}
	return nil
}
