package repository

import (
	"errors"
	"fmt"
	"vmuc-fintech-backend-web-go/domain"

	"gorm.io/gorm"
)

type posgreTrialBalanceRepository struct {
	DB *gorm.DB
}

func NewPostgreTrialBalance(client *gorm.DB) domain.TrialBalanceRepository {
	return &posgreTrialBalanceRepository{
		DB: client,
	}
}

func (a *posgreTrialBalanceRepository) RetrieveAllTrialBalance() ([]domain.TrialBalance, error) {
	var res []domain.TrialBalance
	err := a.DB.
		Model(domain.TrialBalance{}).
		Find(&res).Error
	if err != nil {
		return []domain.TrialBalance{}, err
	}
	fmt.Println("retrieve all TrialBalance ", res)
	return res, nil
}

func (a *posgreTrialBalanceRepository) RetrieveTrialBalanceByID(id uint) (*domain.TrialBalance, error) {
	var res domain.TrialBalance
	err := a.DB.
		Model(domain.TrialBalance{}).
		Where("id = ?", id).
		Take(&res).Error
	if err != nil {
		return &domain.TrialBalance{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.TrialBalance{}, fmt.Errorf("record not found")
	}
	fmt.Println("retrieve TrialBalance by id ", res)
	return &res, nil
}

func (a *posgreTrialBalanceRepository) GetTrialBalanceByTrialBalancePeriodeId(id uint) ([]domain.TrialBalance, error) {
	var res []domain.TrialBalance
	err := a.DB.
		Model(domain.TrialBalance{}).
		Where("id_periode = ?", id).
		Find(&res).Error
	if err != nil {
		return []domain.TrialBalance{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []domain.TrialBalance{}, fmt.Errorf("record not found")
	}
	fmt.Println("retrieve TrialBalance by id ", res)
	return res, nil
}

func (a *posgreTrialBalanceRepository) CreateTrialBalance(TrialBalance *domain.TrialBalance) (*domain.TrialBalance, error) {
	err := a.DB.
		Model(domain.TrialBalance{}).
		Create(TrialBalance).Error
	if err != nil {
		return &domain.TrialBalance{}, err
	}
	fmt.Println("create TrialBalance ", TrialBalance)
	return TrialBalance, nil
}

func (a *posgreTrialBalanceRepository) CreateBulkTrialBalance(TrialBalances []*domain.TrialBalance) ([]*domain.TrialBalance, error) {
	err := a.DB.
		Model(domain.TrialBalance{}).
		Create(&TrialBalances).Error
	if err != nil {
		return []*domain.TrialBalance{}, err
	}
	fmt.Println("create bulk TrialBalance ", TrialBalances)
	return TrialBalances, nil
}

func (a *posgreTrialBalanceRepository) UpdateTrialBalance(TrialBalance *domain.TrialBalance) (*domain.TrialBalance, error) {
	err := a.DB.
		Model(domain.TrialBalance{}).
		Where("id = ?", TrialBalance.ID).
		Updates(TrialBalance).Error
	if err != nil {
		return &domain.TrialBalance{}, err
	}
	fmt.Println("update TrialBalance ", TrialBalance)
	return TrialBalance, nil
}

func (a *posgreTrialBalanceRepository) UpdateBulkTrialBalance(TrialBalances []*domain.TrialBalance) ([]*domain.TrialBalance, error) {
	err := a.DB.
		Model(domain.TrialBalance{}).
		Updates(&TrialBalances).Error
	if err != nil {
		return []*domain.TrialBalance{}, err
	}
	fmt.Println("update bulk TrialBalance ", TrialBalances)
	return TrialBalances, nil
}

func (a *posgreTrialBalanceRepository) DeleteTrialBalance(id uint) error {
	err := a.DB.
		Model(domain.TrialBalance{}).
		Where("id = ?", id).
		Delete(&domain.TrialBalance{}).Error
	if err != nil {
		return err
	}
	fmt.Println("delete TrialBalance ", id)
	return nil
}
