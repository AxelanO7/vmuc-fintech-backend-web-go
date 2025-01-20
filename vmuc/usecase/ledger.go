package usecase

import (
	"context"
	"time"
	"vmuc-fintech-backend-web-go/domain"
)

type generalLedgerUseCase struct {
	generalLedgerRepository  domain.GeneralLedgerRepository
	periodeRepository        domain.PeriodeRepository
	generalJournalRepository domain.GeneralJournalRepository
	contextTimeout           time.Duration
}

func NewGeneralLedgerUseCase(generalLedger domain.GeneralLedgerRepository, periode domain.PeriodeRepository, generalJournal domain.GeneralJournalRepository, t time.Duration) domain.GeneralLedgerUseCase {
	return &generalLedgerUseCase{
		generalLedgerRepository:  generalLedger,
		periodeRepository:        periode,
		generalJournalRepository: generalJournal,
		contextTimeout:           t,
	}
}

func (c *generalLedgerUseCase) FetchGeneralLedgerByID(ctx context.Context, id uint, opt bool) (map[string]any, error) {

	res, err := c.generalLedgerRepository.RetrieveGeneralLedgerByID(id)
	if err != nil {
		return nil, err
	}

	if opt {
		resPeriode, err := c.periodeRepository.GetBeriodeByPeriode(res.Date)
		if err != nil {
			return nil, err
		}
		resGeneralJournal, err := c.generalLedgerRepository.GetGeneralLedgerByGeneralLedgerPeriodeId(resPeriode.ID)
		if err != nil {
			return nil, err
		}

		payload := map[string]any{
			"general_ledger":  res,
			"general_journal": resGeneralJournal,
		}

		return payload, nil
	}

	payload := map[string]any{
		"general_ledger": res,
	}

	return payload, nil
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
