package repository

import (
	"assyarif-backend-web-go/domain"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type posgreOutletRepository struct {
	DB *gorm.DB
}

func NewPostgreOutlet(client *gorm.DB) domain.OutletRepository {
	return &posgreOutletRepository{
		DB: client,
	}
}

func (a *posgreOutletRepository) RetrieveAllOutlet() ([]domain.Outlet, error) {
	var res []domain.Outlet
	err := a.DB.
		Model(domain.Outlet{}).
		Preload("User").
		Find(&res).Error
	if err != nil {
		return []domain.Outlet{}, err
	}
	fmt.Println(res)
	return res, nil
}

func (a *posgreOutletRepository) RetrieveOutletByID(id uint) (*domain.Outlet, error) {
	var res domain.Outlet
	err := a.DB.
		Model(domain.Outlet{}).
		Where("id = ?", id).
		Take(&res).Error
	if err != nil {
		return &domain.Outlet{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.Outlet{}, fmt.Errorf("record not found")
	}
	fmt.Println(res)
	return &res, nil
}

func (a *posgreOutletRepository) CreateOutlet(user *domain.Outlet) (*domain.Outlet, error) {
	err := a.DB.
		Model(domain.Outlet{}).
		Create(user).Error
	if err != nil {
		return &domain.Outlet{}, err
	}
	fmt.Println(user)
	return user, nil
}

func (a *posgreOutletRepository) UpdateOutlet(user *domain.Outlet) (*domain.Outlet, error) {
	err := a.DB.
		Model(domain.Outlet{}).
		Where("id = ?", user.ID).
		Updates(user).Error
	if err != nil {
		return &domain.Outlet{}, err
	}
	fmt.Println(user)
	return user, nil
}

func (a *posgreOutletRepository) DeleteOutlet(id uint) error {
	err := a.DB.
		Model(domain.Outlet{}).
		Where("id = ?", id).
		Delete(&domain.Outlet{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *posgreOutletRepository) ShowOutletByIDUser(id uint) (*domain.Outlet, error) {
	var res domain.Outlet
	err := a.DB.
		Model(domain.Outlet{}).
		Where("id_user = ?", id).
		Take(&res).Error
	if err != nil {
		return &domain.Outlet{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.Outlet{}, fmt.Errorf("record not found")
	}
	fmt.Println(res)
	return &res, nil
}
