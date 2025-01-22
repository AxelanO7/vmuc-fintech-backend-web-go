package usecase

import (
	"context"
	"time"
	"vmuc-fintech-backend-web-go/domain"
)

type worksheetUseCase struct {
	worksheetRepository       domain.WorksheetRepository
	generalLedgerRepository   domain.GeneralJournalRepository
	adjusmentEntriesRepostory domain.AdjusmentEntriesRepository
	periodeRepository         domain.PeriodeRepository
	contextTimeout            time.Duration
}

func NewWorksheetUseCase(worksheet domain.WorksheetRepository, gl domain.GeneralJournalRepository, ae domain.AdjusmentEntriesRepository, per domain.PeriodeRepository, t time.Duration) domain.WorksheetUseCase {
	return &worksheetUseCase{
		worksheetRepository:       worksheet,
		generalLedgerRepository:   gl,
		adjusmentEntriesRepostory: ae,
		periodeRepository:         per,
		contextTimeout:            t,
	}
}

func (c *worksheetUseCase) FetchWorksheetByID(ctx context.Context, id uint, opt bool) (map[string]any, error) {
	res, err := c.worksheetRepository.RetrieveWorksheetByID(id)
	if err != nil {
		return nil, err
	}

	if opt {
		resPeriode, err := c.periodeRepository.GetBeriodeByPeriode(res.Date)
		if err != nil {
			return nil, err
		}
		resGeneralJournal, err := c.generalLedgerRepository.GetGeneralJournalByGeneralJournalPeriodeId(resPeriode.ID)
		if err != nil {
			return nil, err
		}

		resAdjusmentEntries, err := c.adjusmentEntriesRepostory.GetAdjusmentEntriesByAdjusmentEntriesPeriodeId(resPeriode.ID)
		if err != nil {
			return nil, err
		}

		payloadArray := []map[string]any{}

		for _, generalJournal := range resGeneralJournal {
			for _, adjusmentEntry := range resAdjusmentEntries {

				if generalJournal.NameAccount == adjusmentEntry.NameAccount {
					var kredit, debit float64

					if generalJournal.Kredit != 0 && adjusmentEntry.Debit != 0 {

						result := generalJournal.Kredit - adjusmentEntry.Debit
						if result < 0 {

							kredit = -result
						} else {

							debit = result
						}
					} else if generalJournal.Debit != 0 && adjusmentEntry.Kredit != 0 {

						result := generalJournal.Debit - adjusmentEntry.Kredit
						if result < 0 {

							kredit = -result
						} else {

							debit = result
						}
					} else if generalJournal.Debit != 0 && adjusmentEntry.Debit != 0 {

						debit = generalJournal.Debit + adjusmentEntry.Debit
					} else if generalJournal.Kredit != 0 && adjusmentEntry.Kredit != 0 {

						kredit = generalJournal.Kredit + adjusmentEntry.Kredit
					}

					payload := map[string]any{
						"name_account": generalJournal.NameAccount,
						"kredit":       kredit,
						"debit":        debit,
					}

					payloadArray = append(payloadArray, payload)
				}
			}
		}

		payloadResult := map[string]any{
			"worksheet":      res,
			"data_worksheet": payloadArray,
		}

		return payloadResult, nil

	}

	payload := map[string]any{
		"worksheet": res,
	}

	return payload, nil
}

func (c *worksheetUseCase) FetchWorksheets(ctx context.Context) ([]domain.Worksheet, error) {
	res, err := c.worksheetRepository.RetrieveAllWorksheet()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *worksheetUseCase) AddWorksheet(ctx context.Context, req *domain.Worksheet) (*domain.Worksheet, error) {
	res, err := c.worksheetRepository.CreateWorksheet(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *worksheetUseCase) AddBulkWorksheet(ctx context.Context, req []*domain.Worksheet) ([]*domain.Worksheet, error) {
	res, err := c.worksheetRepository.CreateBulkWorksheet(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *worksheetUseCase) EditWorksheet(ctx context.Context, req *domain.Worksheet) (*domain.Worksheet, error) {
	res, err := c.worksheetRepository.UpdateWorksheet(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *worksheetUseCase) EditBulkWorksheet(ctx context.Context, req []*domain.Worksheet) ([]*domain.Worksheet, error) {
	res, err := c.worksheetRepository.UpdateBulkWorksheet(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *worksheetUseCase) DeleteWorksheet(ctx context.Context, id uint) error {
	err := c.worksheetRepository.DeleteWorksheet(id)
	if err != nil {
		return err
	}
	return nil
}
