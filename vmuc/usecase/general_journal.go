package usecase

import (
	"context"
	"fmt"
	"time"
	"vmuc-fintech-backend-web-go/domain"
)

type generalJournalUseCase struct {
	generalJournalRepository domain.GeneralJournalRepository
	contextTimeout           time.Duration
}

func NewGeneralJournalUseCase(generalJournal domain.GeneralJournalRepository, t time.Duration) domain.GeneralJournalUseCase {
	return &generalJournalUseCase{
		generalJournalRepository: generalJournal,
		contextTimeout:           t,
	}
}

func (c *generalJournalUseCase) FetchGeneralJournalByID(ctx context.Context, id uint) (*domain.GeneralJournal, error) {
	res, err := c.generalJournalRepository.RetrieveGeneralJournalByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *generalJournalUseCase) FetchGeneralJournals(ctx context.Context) ([]domain.GeneralJournal, error) {
	res, err := c.generalJournalRepository.RetrieveAllGeneralJournal()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *generalJournalUseCase) AddGeneralJournal(ctx context.Context, req *domain.GeneralJournal) (*domain.GeneralJournal, error) {
	res, err := c.generalJournalRepository.CreateGeneralJournal(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *generalJournalUseCase) AddBulkGeneralJournal(ctx context.Context, req []*domain.GeneralJournal) ([]*domain.GeneralJournal, error) {
	res, err := c.generalJournalRepository.CreateBulkGeneralJournal(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *generalJournalUseCase) EditGeneralJournal(ctx context.Context, req *domain.GeneralJournal) (*domain.GeneralJournal, error) {
	res, err := c.generalJournalRepository.UpdateGeneralJournal(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *generalJournalUseCase) EditBulkGeneralJournal(ctx context.Context, req []*domain.GeneralJournal) ([]*domain.GeneralJournal, error) {
	fmt.Println(&req)
	res, err := c.generalJournalRepository.UpdateBulkGeneralJournal(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *generalJournalUseCase) DeleteGeneralJournal(ctx context.Context, id uint) error {
	err := c.generalJournalRepository.DeleteGeneralJournal(id)
	if err != nil {
		return err
	}
	return nil
}
