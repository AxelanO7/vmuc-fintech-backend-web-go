package repository

import (
	"errors"
	"fmt"
	"vmuc-fintech-backend-web-go/domain"

	"gorm.io/gorm"
)

type posgreAdjusmentEntriesRepository struct {
	DB *gorm.DB
}

func NewPostgreAdjusmentEntries(client *gorm.DB) domain.AdjusmentEntriesRepository {
	return &posgreAdjusmentEntriesRepository{
		DB: client,
	}
}

func (a *posgreAdjusmentEntriesRepository) RetrieveAllAdjusmentEntries() ([]domain.AdjusmentEntries, error) {
	var res []domain.AdjusmentEntries
	err := a.DB.
		Model(domain.AdjusmentEntries{}).
		Find(&res).Error
	if err != nil {
		return []domain.AdjusmentEntries{}, err
	}
	fmt.Println("retrieve all AdjusmentEntries ", res)
	return res, nil
}

func (a *posgreAdjusmentEntriesRepository) RetrieveAdjusmentEntriesByID(id uint) (*domain.AdjusmentEntries, error) {
	var res domain.AdjusmentEntries
	err := a.DB.
		Model(domain.AdjusmentEntries{}).
		Where("id = ?", id).
		Take(&res).Error
	if err != nil {
		return &domain.AdjusmentEntries{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.AdjusmentEntries{}, fmt.Errorf("record not found")
	}
	fmt.Println("retrieve AdjusmentEntries by id ", res)
	return &res, nil
}

func (a *posgreAdjusmentEntriesRepository) GetAdjusmentEntriesByAdjusmentEntriesPeriodeId(id uint) ([]domain.AdjusmentEntries, error) {
	var res []domain.AdjusmentEntries
	err := a.DB.
		Model(domain.AdjusmentEntries{}).
		Where("id_periode = ?", id).
		Find(&res).Error
	if err != nil {
		return []domain.AdjusmentEntries{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []domain.AdjusmentEntries{}, fmt.Errorf("record not found")
	}
	fmt.Println("retrieve AdjusmentEntries by id ", res)
	return res, nil
}

func (a *posgreAdjusmentEntriesRepository) CreateAdjusmentEntries(AdjusmentEntries *domain.AdjusmentEntries) (*domain.AdjusmentEntries, error) {
	err := a.DB.
		Model(domain.AdjusmentEntries{}).
		Create(AdjusmentEntries).Error
	if err != nil {
		return &domain.AdjusmentEntries{}, err
	}
	fmt.Println("create AdjusmentEntries ", AdjusmentEntries)
	return AdjusmentEntries, nil
}

func (a *posgreAdjusmentEntriesRepository) CreateBulkAdjusmentEntries(AdjusmentEntriess []*domain.AdjusmentEntries) ([]*domain.AdjusmentEntries, error) {
	err := a.DB.
		Model(domain.AdjusmentEntries{}).
		Create(&AdjusmentEntriess).Error
	if err != nil {
		return []*domain.AdjusmentEntries{}, err
	}
	fmt.Println("create bulk AdjusmentEntries ", AdjusmentEntriess)
	return AdjusmentEntriess, nil
}

func (a *posgreAdjusmentEntriesRepository) UpdateAdjusmentEntries(AdjusmentEntries *domain.AdjusmentEntries) (*domain.AdjusmentEntries, error) {
	err := a.DB.
		Model(domain.AdjusmentEntries{}).
		Where("id = ?", AdjusmentEntries.ID).
		Updates(AdjusmentEntries).Error
	if err != nil {
		return &domain.AdjusmentEntries{}, err
	}
	fmt.Println("update AdjusmentEntries ", AdjusmentEntries)
	return AdjusmentEntries, nil
}

func (a *posgreAdjusmentEntriesRepository) UpdateBulkAdjusmentEntries(AdjusmentEntriess []*domain.AdjusmentEntries) ([]*domain.AdjusmentEntries, error) {
	err := a.DB.
		Model(domain.AdjusmentEntries{}).
		Updates(&AdjusmentEntriess).Error
	if err != nil {
		return []*domain.AdjusmentEntries{}, err
	}
	fmt.Println("update bulk AdjusmentEntries ", AdjusmentEntriess)
	return AdjusmentEntriess, nil
}

func (a *posgreAdjusmentEntriesRepository) DeleteAdjusmentEntries(id uint) error {
	err := a.DB.
		Model(domain.AdjusmentEntries{}).
		Where("id = ?", id).
		Delete(&domain.AdjusmentEntries{}).Error
	if err != nil {
		return err
	}
	fmt.Println("delete AdjusmentEntries ", id)
	return nil
}
