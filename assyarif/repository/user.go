package repository

import (
	"assyarif-backend-web-go/domain"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type posgreUserRepository struct {
	DB *gorm.DB
}

func NewPostgreUser(client *gorm.DB) domain.UserRepository {
	return &posgreUserRepository{
		DB: client,
	}
}

func (a *posgreUserRepository) GetUser(username string) (*domain.User, error) {
	var res domain.User
	err := a.DB.
		Model(domain.User{}).
		Where("username = ?", username).
		Take(&res).Error
	if err != nil {
		return &domain.User{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.User{}, fmt.Errorf("record not found")
	}
	fmt.Println(res)
	return &res, nil
}

func (a *posgreUserRepository) GetUserById(id uint) (*domain.User, error) {
	var res domain.User
	err := a.DB.
		Model(domain.User{}).
		Where("id = ?", id).
		Take(&res).Error
	if err != nil {
		return &domain.User{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.User{}, fmt.Errorf("record not found")
	}
	fmt.Println(res)
	return &res, nil
}
