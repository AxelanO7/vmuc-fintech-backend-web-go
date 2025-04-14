package repository

import (
	"errors"
	"fmt"
	"vmuc-fintech-backend-web-go/domain"

	"gorm.io/gorm"
)

type posgreGeneralJournalRepository struct {
	DB *gorm.DB
}

func NewPostgreGeneralJournal(client *gorm.DB) domain.GeneralJournalRepository {
	return &posgreGeneralJournalRepository{
		DB: client,
	}
}

func (a *posgreGeneralJournalRepository) RetrieveAllGeneralJournal() ([]domain.GeneralJournal, error) {
	var res []domain.GeneralJournal
	err := a.DB.
		Model(domain.GeneralJournal{}).
		Find(&res).Error
	if err != nil {
		return []domain.GeneralJournal{}, err
	}
	fmt.Println("retrieve all GeneralJournal ", res)
	return res, nil
}

func (a *posgreGeneralJournalRepository) RetrieveGeneralJournalByID(id uint) (*domain.GeneralJournal, error) {
	var res domain.GeneralJournal
	err := a.DB.
		Model(domain.GeneralJournal{}).
		Where("id = ?", id).
		Take(&res).Error
	if err != nil {
		return &domain.GeneralJournal{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.GeneralJournal{}, fmt.Errorf("record not found")
	}
	fmt.Println("retrieve GeneralJournal by id ", res)
	return &res, nil
}

func (a *posgreGeneralJournalRepository) GetGeneralJournalByGeneralJournalPeriodeId(id uint) ([]domain.GeneralJournal, error) {
	var res []domain.GeneralJournal
	err := a.DB.
		Model(domain.GeneralJournal{}).
		Where("id_periode = ?", id).
		Find(&res).Error
	if err != nil {
		return []domain.GeneralJournal{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []domain.GeneralJournal{}, fmt.Errorf("record not found")
	}
	fmt.Println("retrieve GeneralJournal by id ", res)
	return res, nil
}

func (a *posgreGeneralJournalRepository) CreateGeneralJournal(GeneralJournal *domain.GeneralJournal) (*domain.GeneralJournal, error) {
	err := a.DB.
		Model(domain.GeneralJournal{}).
		Create(GeneralJournal).Error
	if err != nil {
		return &domain.GeneralJournal{}, err
	}
	fmt.Println("create GeneralJournal ", GeneralJournal)
	return GeneralJournal, nil
}

func (a *posgreGeneralJournalRepository) CreateBulkGeneralJournal(GeneralJournals []*domain.GeneralJournal) ([]*domain.GeneralJournal, error) {
	err := a.DB.
		Model(domain.GeneralJournal{}).
		Create(&GeneralJournals).Error
	if err != nil {
		return []*domain.GeneralJournal{}, err
	}
	fmt.Println("create bulk GeneralJournal ", GeneralJournals)
	return GeneralJournals, nil
}

func (a *posgreGeneralJournalRepository) UpdateGeneralJournal(GeneralJournal *domain.GeneralJournal) (*domain.GeneralJournal, error) {
	err := a.DB.
		Model(domain.GeneralJournal{}).
		Where("id = ?", GeneralJournal.ID).
		Updates(GeneralJournal).Error
	if err != nil {
		return &domain.GeneralJournal{}, err
	}
	fmt.Println("update GeneralJournal ", GeneralJournal)
	return GeneralJournal, nil
}

func (a *posgreGeneralJournalRepository) UpdateBulkGeneralJournal(GeneralJournals []*domain.GeneralJournal) ([]*domain.GeneralJournal, error) {
	for _, val := range GeneralJournals {
		err := a.DB.
			Model(domain.GeneralJournal{}).
			Where("id = ?", val.ID).
			Updates(val).Error
		if err != nil {
			fmt.Println("update bulk GeneralJournal ", GeneralJournals)
			return []*domain.GeneralJournal{}, err
		}
	}
	return GeneralJournals, nil
}

func (a *posgreGeneralJournalRepository) DeleteGeneralJournal(id uint) error {
	err := a.DB.
		Model(domain.GeneralJournal{}).
		Where("id = ?", id).
		Delete(&domain.GeneralJournal{}).Error
	if err != nil {
		return err
	}
	fmt.Println("delete GeneralJournal ", id)
	return nil
}
