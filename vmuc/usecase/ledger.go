package usecase

import (
	"context"
	"time"
	"vmuc-fintech-backend-web-go/domain"
)

type ledgerUseCase struct {
	ledgerRepository         domain.LedgerRepository
	periodeRepository        domain.PeriodeRepository
	generalJournalRepository domain.GeneralJournalRepository
	contextTimeout           time.Duration
}

func NewLedgerUseCase(generalLedger domain.LedgerRepository, periode domain.PeriodeRepository, generalJournal domain.GeneralJournalRepository, t time.Duration) domain.LedgerUseCase {
	return &ledgerUseCase{
		ledgerRepository:         generalLedger,
		periodeRepository:        periode,
		generalJournalRepository: generalJournal,
		contextTimeout:           t,
	}
}

func (c *ledgerUseCase) FetchLedgerByID(ctx context.Context, id uint, opt bool) (map[string]any, error) {

	res, err := c.ledgerRepository.RetrieveLedgerByID(id)
	if err != nil {
		return nil, err
	}

	if opt {
		resPeriode, err := c.periodeRepository.GetPeriodeByPeriode(res.Date)
		if err != nil {
			return nil, err
		}
		resGeneralJournal, err := c.generalJournalRepository.GetGeneralJournalByGeneralJournalPeriodeId(resPeriode.ID)
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

func (c *ledgerUseCase) FetchLedgers(ctx context.Context) ([]domain.Ledger, error) {
	res, err := c.ledgerRepository.RetrieveLedgers()
	if err != nil {
		return nil, err
	}
	for i := range res {
		resPeriode, err := c.periodeRepository.GetPeriodeByPeriode(res[i].Date)
		if err != nil {
			return nil, err
		}
		resGeneralJournal, err := c.generalJournalRepository.GetGeneralJournalByGeneralJournalPeriodeId(resPeriode.ID)
		if err != nil {
			return nil, err
		}

		res[i].GeneralJournal = resGeneralJournal
	}
	return res, nil
}

func (c *ledgerUseCase) AddLedger(ctx context.Context, req *domain.Ledger) (*domain.Ledger, error) {
	res, err := c.ledgerRepository.CreateLedger(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *ledgerUseCase) AddBulkLedger(ctx context.Context, req []*domain.Ledger) ([]*domain.Ledger, error) {
	res, err := c.ledgerRepository.CreateBulkLedger(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *ledgerUseCase) EditLedger(ctx context.Context, req *domain.Ledger) (*domain.Ledger, error) {
	res, err := c.ledgerRepository.UpdateLedger(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *ledgerUseCase) EditBulkLedger(ctx context.Context, req []*domain.Ledger) ([]*domain.Ledger, error) {
	res, err := c.ledgerRepository.UpdateBulkLedger(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *ledgerUseCase) DeleteLedger(ctx context.Context, id uint) error {
	err := c.ledgerRepository.DeleteLedger(id)
	if err != nil {
		return err
	}
	return nil
}
