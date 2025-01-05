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
	res := []domain.Ref{}
	err := a.DB.
		Model(domain.Ref{}).
		Find(&res).Error
	if err != nil {
		return []domain.Ref{}, err
	}
	if len(res) == 0 {
		return []domain.Ref{}, fmt.Errorf("record not found")
	}
	fmt.Println("retrieve refs ", res)
	return res, nil
}

func (a *posgreRefRepository) RetrieveRefByID(id uint) (*domain.Ref, error) {
	res := domain.Ref{}
	err := a.DB.
		Model(domain.Ref{}).
		Where("id = ?", id).
		Take(&res).Error
	if err != nil {
		return &domain.Ref{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.Ref{}, fmt.Errorf("record not found")
	}
	fmt.Println("retrieve ref by id ", res)
	return &res, nil
}

func (a *posgreRefRepository) CreateRef(ref *domain.Ref) (*domain.Ref, error) {
	err := a.DB.
		Model(domain.Ref{}).
		Create(ref).Error
	if err != nil {
		return &domain.Ref{}, err
	}
	fmt.Println("create ref ", ref)
	return ref, nil
}

func (a *posgreRefRepository) UpdateRef(ref *domain.Ref) (*domain.Ref, error) {
	err := a.DB.
		Model(domain.Ref{}).
		Where("id = ?", ref.ID).
		Updates(ref).Error
	if err != nil {
		return &domain.Ref{}, err
	}
	fmt.Println("update ref ", ref)
	return ref, nil
}

func (a *posgreRefRepository) DeleteRef(id uint) error {
	err := a.DB.
		Model(domain.Ref{}).
		Where("id = ?", id).
		Delete(&domain.Ref{}).Error
	if err != nil {
		return err
	}
	fmt.Println("delete ref ", id)
	return nil
}
