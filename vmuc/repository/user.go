package repository

import (
	"errors"
	"fmt"
	"vmuc-fintech-backend-web-go/domain"

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

func (a *posgreUserRepository) RetrieveAllUser() ([]domain.User, error) {
	var res []domain.User
	err := a.DB.
		Model(domain.User{}).
		Find(&res).Error
	if err != nil {
		return []domain.User{}, err
	}
	fmt.Println(" retrieve all user ", res)
	return res, nil
}

func (a *posgreUserRepository) RetrieveByUsername(username string) (*domain.User, error) {
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
	fmt.Println("retrieve user by username ", res)
	return &res, nil
}

func (a *posgreUserRepository) RetrieveUserByID(id uint) (*domain.User, error) {
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
	fmt.Println("retrieve user by id ", res)
	return &res, nil
}

func (a *posgreUserRepository) CreateUser(user *domain.User) (*domain.User, error) {
	err := a.DB.
		Model(domain.User{}).
		Create(user).Error
	if err != nil {
		return &domain.User{}, err
	}
	fmt.Println("create user ", user)
	return user, nil
}

func (a *posgreUserRepository) UpdateUser(user *domain.User) (*domain.User, error) {
	err := a.DB.
		Model(domain.User{}).
		Where("id = ?", user.ID).
		Updates(user).Error
	if err != nil {
		return &domain.User{}, err
	}
	fmt.Println("update user ", user)
	return user, nil
}

func (a *posgreUserRepository) DeleteUser(id uint) error {
	err := a.DB.
		Model(domain.User{}).
		Where("id = ?", id).
		Delete(&domain.User{}).Error
	if err != nil {
		return err
	}
	fmt.Println("delete user ", id)
	return nil
}
