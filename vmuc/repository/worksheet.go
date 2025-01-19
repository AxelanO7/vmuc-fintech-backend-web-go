package repository

import (
	"errors"
	"fmt"
	"vmuc-fintech-backend-web-go/domain"

	"gorm.io/gorm"
)

type posgreWorksheetRepository struct {
	DB *gorm.DB
}

func NewPostgreWorksheet(client *gorm.DB) domain.WorksheetRepository {
	return &posgreWorksheetRepository{
		DB: client,
	}
}

func (a *posgreWorksheetRepository) RetrieveAllWorksheet() ([]domain.Worksheet, error) {
	var res []domain.Worksheet
	err := a.DB.
		Model(domain.Worksheet{}).
		Find(&res).Error
	if err != nil {
		return []domain.Worksheet{}, err
	}
	fmt.Println("retrieve all Worksheet ", res)
	return res, nil
}

func (a *posgreWorksheetRepository) RetrieveWorksheetByID(id uint) (*domain.Worksheet, error) {
	var res domain.Worksheet
	err := a.DB.
		Model(domain.Worksheet{}).
		Where("id = ?", id).
		Take(&res).Error
	if err != nil {
		return &domain.Worksheet{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.Worksheet{}, fmt.Errorf("record not found")
	}
	fmt.Println("retrieve Worksheet by id ", res)
	return &res, nil
}

func (a *posgreWorksheetRepository) GetWorksheetByWorksheetPeriodeId(id uint) ([]domain.Worksheet, error) {
	var res []domain.Worksheet
	err := a.DB.
		Model(domain.Worksheet{}).
		Where("id_periode = ?", id).
		Find(&res).Error
	if err != nil {
		return []domain.Worksheet{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []domain.Worksheet{}, fmt.Errorf("record not found")
	}
	fmt.Println("retrieve Worksheet by id ", res)
	return res, nil
}

func (a *posgreWorksheetRepository) CreateWorksheet(Worksheet *domain.Worksheet) (*domain.Worksheet, error) {
	err := a.DB.
		Model(domain.Worksheet{}).
		Create(Worksheet).Error
	if err != nil {
		return &domain.Worksheet{}, err
	}
	fmt.Println("create Worksheet ", Worksheet)
	return Worksheet, nil
}

func (a *posgreWorksheetRepository) CreateBulkWorksheet(Worksheets []*domain.Worksheet) ([]*domain.Worksheet, error) {
	err := a.DB.
		Model(domain.Worksheet{}).
		Create(&Worksheets).Error
	if err != nil {
		return []*domain.Worksheet{}, err
	}
	fmt.Println("create bulk Worksheet ", Worksheets)
	return Worksheets, nil
}

func (a *posgreWorksheetRepository) UpdateWorksheet(Worksheet *domain.Worksheet) (*domain.Worksheet, error) {
	err := a.DB.
		Model(domain.Worksheet{}).
		Where("id = ?", Worksheet.ID).
		Updates(Worksheet).Error
	if err != nil {
		return &domain.Worksheet{}, err
	}
	fmt.Println("update Worksheet ", Worksheet)
	return Worksheet, nil
}

func (a *posgreWorksheetRepository) UpdateBulkWorksheet(Worksheets []*domain.Worksheet) ([]*domain.Worksheet, error) {
	err := a.DB.
		Model(domain.Worksheet{}).
		Updates(&Worksheets).Error
	if err != nil {
		return []*domain.Worksheet{}, err
	}
	fmt.Println("update bulk Worksheet ", Worksheets)
	return Worksheets, nil
}

func (a *posgreWorksheetRepository) DeleteWorksheet(id uint) error {
	err := a.DB.
		Model(domain.Worksheet{}).
		Where("id = ?", id).
		Delete(&domain.Worksheet{}).Error
	if err != nil {
		return err
	}
	fmt.Println("delete Worksheet ", id)
	return nil
}
