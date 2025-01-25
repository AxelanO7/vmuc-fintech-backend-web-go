package repository

import (
	"errors"
	"fmt"
	"vmuc-fintech-backend-web-go/domain"

	"gorm.io/gorm"
)

type posgreGeneralLedgerRepository struct {
	DB *gorm.DB
}

func NewPostgreGeneralLedger(client *gorm.DB) domain.GeneralLedgerRepository {
	return &posgreGeneralLedgerRepository{
		DB: client,
	}
}

func (a *posgreGeneralLedgerRepository) RetrieveAllGeneralLedger() ([]domain.GeneralLedger, error) {
	var res []domain.GeneralLedger
	err := a.DB.
		Model(domain.GeneralLedger{}).
		Find(&res).Error
	if err != nil {
		return []domain.GeneralLedger{}, err
	}
	fmt.Println("retrieve all GeneralLedger ", res)
	return res, nil
}

func (a *posgreGeneralLedgerRepository) RetrieveGeneralLedgerByID(id uint) (*domain.GeneralLedger, error) {
	var res domain.GeneralLedger
	err := a.DB.
		Model(domain.GeneralLedger{}).
		Where("id = ?", id).
		Take(&res).Error
	if err != nil {
		return &domain.GeneralLedger{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.GeneralLedger{}, fmt.Errorf("record not found")
	}
	fmt.Println("retrieve GeneralLedger by id ", res)
	return &res, nil
}

func (a *posgreGeneralLedgerRepository) GetGeneralLedgerByGeneralLedgerPeriodeId(id uint) ([]domain.GeneralLedger, error) {
	var res []domain.GeneralLedger
	err := a.DB.
		Model(domain.GeneralLedger{}).
		Where("id_periode = ?", id).
		Find(&res).Error
	if err != nil {
		return []domain.GeneralLedger{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []domain.GeneralLedger{}, fmt.Errorf("record not found")
	}
	fmt.Println("retrieve GeneralLedger by id ", res)
	return res, nil
}

func (a *posgreGeneralLedgerRepository) CreateGeneralLedger(GeneralLedger *domain.GeneralLedger) (*domain.GeneralLedger, error) {
	err := a.DB.
		Model(domain.GeneralLedger{}).
		Create(GeneralLedger).Error
	if err != nil {
		return &domain.GeneralLedger{}, err
	}
	fmt.Println("create GeneralLedger ", GeneralLedger)
	return GeneralLedger, nil
}

func (a *posgreGeneralLedgerRepository) CreateBulkGeneralLedger(GeneralLedgers []*domain.GeneralLedger) ([]*domain.GeneralLedger, error) {
	err := a.DB.
		Model(domain.GeneralLedger{}).
		Create(&GeneralLedgers).Error
	if err != nil {
		return []*domain.GeneralLedger{}, err
	}
	fmt.Println("create bulk GeneralLedger ", GeneralLedgers)
	return GeneralLedgers, nil
}

func (a *posgreGeneralLedgerRepository) UpdateGeneralLedger(GeneralLedger *domain.GeneralLedger) (*domain.GeneralLedger, error) {
	err := a.DB.
		Model(domain.GeneralLedger{}).
		Where("id = ?", GeneralLedger.ID).
		Updates(GeneralLedger).Error
	if err != nil {
		return &domain.GeneralLedger{}, err
	}
	fmt.Println("update GeneralLedger ", GeneralLedger)
	return GeneralLedger, nil
}

func (a *posgreGeneralLedgerRepository) UpdateBulkGeneralLedger(GeneralLedgers []*domain.GeneralLedger) ([]*domain.GeneralLedger, error) {
	err := a.DB.
		Model(domain.GeneralLedger{}).
		Updates(&GeneralLedgers).Error
	if err != nil {
		return []*domain.GeneralLedger{}, err
	}
	fmt.Println("update bulk GeneralLedger ", GeneralLedgers)
	return GeneralLedgers, nil
}

func (a *posgreGeneralLedgerRepository) DeleteGeneralLedger(id uint) error {
	err := a.DB.
		Model(domain.GeneralLedger{}).
		Where("id = ?", id).
		Delete(&domain.GeneralLedger{}).Error
	if err != nil {
		return err
	}
	fmt.Println("delete GeneralLedger ", id)
	return nil
}
