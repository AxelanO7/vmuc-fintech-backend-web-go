package repository

import (
	"errors"
	"fmt"
	"vmuc-fintech-backend-web-go/domain"

	"gorm.io/gorm"
)

type posgreRefRepository struct {
	DB *gorm.DB
}

func NewPostgreRef(client *gorm.DB) domain.RefRepository {
	return &posgreRefRepository{
		DB: client,
	}
}

func (a *posgreRefRepository) RetrieveRefs() ([]domain.Ref, error) {
	var res []domain.Ref
	err := a.DB.
		Model(domain.Ref{}).
		Preload("User").
		Find(&res).Error
	if err != nil {
		return []domain.Ref{}, err
	}
	fmt.Println(res)
	return res, nil
}

func (a *posgreRefRepository) RetrieveRefByID(id uint) (*domain.Ref, error) {
	var res domain.Ref
	err := a.DB.
		Model(domain.Ref{}).
		Where("id = ?", id).
		Preload("User").
		Take(&res).Error
	if err != nil {
		return &domain.Ref{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.Ref{}, fmt.Errorf("record not found")
	}
	fmt.Println(res)
	return &res, nil
}

func (a *posgreRefRepository) CreateRef(user *domain.Ref) (*domain.Ref, error) {
	err := a.DB.
		Model(domain.Ref{}).
		Create(user).Error
	if err != nil {
		return &domain.Ref{}, err
	}
	fmt.Println(user)
	return user, nil
}

func (a *posgreRefRepository) UpdateRef(user *domain.Ref) (*domain.Ref, error) {
	err := a.DB.
		Model(domain.Ref{}).
		Where("id = ?", user.ID).
		Updates(user).Error
	if err != nil {
		return &domain.Ref{}, err
	}
	fmt.Println(user)
	return user, nil
}

func (a *posgreRefRepository) DeleteRef(id uint) error {
	err := a.DB.
		Model(domain.Ref{}).
		Where("id = ?", id).
		Delete(&domain.Ref{}).Error
	if err != nil {
		return err
	}
	return nil
}
