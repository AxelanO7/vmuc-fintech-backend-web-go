package usecase

import (
	"context"
	"time"
	"vmuc-fintech-backend-web-go/domain"
)

type generalLedgerUseCase struct {
	generalLedgerRepository domain.GeneralLedgerRepository
	contextTimeout          time.Duration
}

func NewGeneralLedgerUseCase(generalLedger domain.GeneralLedgerRepository, t time.Duration) domain.GeneralLedgerUseCase {
	return &generalLedgerUseCase{
		generalLedgerRepository: generalLedger,
		contextTimeout:          t,
	}
}

func (c *generalLedgerUseCase) FetchGeneralLedgerByID(ctx context.Context, id uint) (*domain.GeneralLedger, error) {
	res, err := c.generalLedgerRepository.RetrieveGeneralLedgerByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *generalLedgerUseCase) FetchGeneralLedgers(ctx context.Context) ([]domain.GeneralLedger, error) {
	res, err := c.generalLedgerRepository.RetrieveAllGeneralLedger()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *generalLedgerUseCase) AddGeneralLedger(ctx context.Context, req *domain.GeneralLedger) (*domain.GeneralLedger, error) {
	res, err := c.generalLedgerRepository.CreateGeneralLedger(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *generalLedgerUseCase) AddBulkGeneralLedger(ctx context.Context, req []*domain.GeneralLedger) ([]*domain.GeneralLedger, error) {
	res, err := c.generalLedgerRepository.CreateBulkGeneralLedger(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *generalLedgerUseCase) EditGeneralLedger(ctx context.Context, req *domain.GeneralLedger) (*domain.GeneralLedger, error) {
	res, err := c.generalLedgerRepository.UpdateGeneralLedger(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *generalLedgerUseCase) EditBulkGeneralLedger(ctx context.Context, req []*domain.GeneralLedger) ([]*domain.GeneralLedger, error) {
	res, err := c.generalLedgerRepository.UpdateBulkGeneralLedger(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *generalLedgerUseCase) DeleteGeneralLedger(ctx context.Context, id uint) error {
	err := c.generalLedgerRepository.DeleteGeneralLedger(id)
	if err != nil {
		return err
	}
	return nil
}
