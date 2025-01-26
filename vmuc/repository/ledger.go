package repository

import (
	"errors"
	"fmt"
	"vmuc-fintech-backend-web-go/domain"

	"gorm.io/gorm"
)

type posgreLedgerRepository struct {
	DB *gorm.DB
}

func NewPostgreLedger(client *gorm.DB) domain.LedgerRepository {
	return &posgreLedgerRepository{
		DB: client,
	}
}

func (a *posgreLedgerRepository) RetrieveLedgers() ([]domain.Ledger, error) {
	var res []domain.Ledger
	err := a.DB.
		Model(domain.Ledger{}).
		Find(&res).Error
	if err != nil {
		return []domain.Ledger{}, err
	}
	fmt.Println("retrieve all GeneralLedger ", res)
	return res, nil
}

func (a *posgreLedgerRepository) RetrieveLedgerByID(id uint) (*domain.Ledger, error) {
	var res domain.Ledger
	err := a.DB.
		Model(domain.Ledger{}).
		Where("id = ?", id).
		Take(&res).Error
	if err != nil {
		return &domain.Ledger{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.Ledger{}, fmt.Errorf("record not found")
	}
	fmt.Println("retrieve GeneralLedger by id ", res)
	return &res, nil
}

func (a *posgreLedgerRepository) GetLedgerByPeriodeId(id uint) ([]domain.Ledger, error) {
	var res []domain.Ledger
	err := a.DB.
		Model(domain.Ledger{}).
		Where("id_periode = ?", id).
		Find(&res).Error
	if err != nil {
		return []domain.Ledger{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []domain.Ledger{}, fmt.Errorf("record not found")
	}
	fmt.Println("retrieve GeneralLedger by id ", res)
	return res, nil
}

func (a *posgreLedgerRepository) CreateLedger(GeneralLedger *domain.Ledger) (*domain.Ledger, error) {
	err := a.DB.
		Model(domain.Ledger{}).
		Create(GeneralLedger).Error
	if err != nil {
		return &domain.Ledger{}, err
	}
	fmt.Println("create GeneralLedger ", GeneralLedger)
	return GeneralLedger, nil
}

func (a *posgreLedgerRepository) CreateBulkLedger(GeneralLedgers []*domain.Ledger) ([]*domain.Ledger, error) {
	err := a.DB.
		Model(domain.Ledger{}).
		Create(&GeneralLedgers).Error
	if err != nil {
		return []*domain.Ledger{}, err
	}
	fmt.Println("create bulk GeneralLedger ", GeneralLedgers)
	return GeneralLedgers, nil
}

func (a *posgreLedgerRepository) UpdateLedger(GeneralLedger *domain.Ledger) (*domain.Ledger, error) {
	err := a.DB.
		Model(domain.Ledger{}).
		Where("id = ?", GeneralLedger.ID).
		Updates(GeneralLedger).Error
	if err != nil {
		return &domain.Ledger{}, err
	}
	fmt.Println("update GeneralLedger ", GeneralLedger)
	return GeneralLedger, nil
}

func (a *posgreLedgerRepository) UpdateBulkLedger(GeneralLedgers []*domain.Ledger) ([]*domain.Ledger, error) {
	err := a.DB.
		Model(domain.Ledger{}).
		Updates(&GeneralLedgers).Error
	if err != nil {
		return []*domain.Ledger{}, err
	}
	fmt.Println("update bulk GeneralLedger ", GeneralLedgers)
	return GeneralLedgers, nil
}

func (a *posgreLedgerRepository) DeleteLedger(id uint) error {
	err := a.DB.
		Model(domain.Ledger{}).
		Where("id = ?", id).
		Delete(&domain.Ledger{}).Error
	if err != nil {
		return err
	}
	fmt.Println("delete GeneralLedger ", id)
	return nil
}
