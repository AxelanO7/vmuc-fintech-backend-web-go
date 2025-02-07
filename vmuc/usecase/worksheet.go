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
		resPeriode, err := c.periodeRepository.GetPeriodeByPeriode(res.Date)
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

		// Array to store the result
		payloadArray := []map[string]any{}

		// Loop through each general journal entry
		for _, generalJournal := range resGeneralJournal {
			// Check if the generalJournal is valid (ensure it has necessary data)
			if generalJournal.NameAccount == "" {
				continue // Skip invalid entries
			}

			// Loop through each adjustment entry
			for _, adjusmentEntry := range resAdjusmentEntries {
				// Ensure the adjustment entry is valid (has necessary data)
				if adjusmentEntry.NameAccount == "" {
					continue // Skip invalid entries
				}

				// Check if the accounts match
				if generalJournal.NameAccount == adjusmentEntry.NameAccount {
					var kredit, debit float64

					// Handle various cases of debit and credit
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

					// Prepare payload
					payload := map[string]any{
						"name_account": generalJournal.NameAccount,
						"kredit":       kredit,
						"debit":        debit,
					}

					// Append the payload to the result array
					payloadArray = append(payloadArray, payload)
				}
			}
		}

		// Combine the worksheet and its related data into the result
		payloadResult := map[string]any{
			"worksheet":      res,
			"data_worksheet": payloadArray,
		}

		return payloadResult, nil
	}

	// If no options are set, return just the worksheet data
	payload := map[string]any{
		"worksheet": res,
	}

	return payload, nil
}

func (c *worksheetUseCase) FetchWorksheets(ctx context.Context, opt bool) ([]map[string]any, error) {
	res, err := c.worksheetRepository.RetrieveAllWorksheet()
	if err != nil {
		return nil, err
	}

	if opt {
		// Array untuk menyimpan hasil akhir
		var payloadArray []map[string]any

		// Loop melalui setiap worksheet yang diperoleh
		for _, worksheet := range res {
			resPeriode, err := c.periodeRepository.GetPeriodeByPeriode(worksheet.Date)
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

			// Array untuk menyimpan hasil pengolahan data worksheet ini
			var worksheetPayloadArray []map[string]any

			// Loop melalui setiap entri jurnal umum
			for _, generalJournal := range resGeneralJournal {
				if generalJournal.NameAccount == "" {
					continue // Skip jika tidak valid
				}

				// Loop melalui setiap entri penyesuaian
				for _, adjusmentEntry := range resAdjusmentEntries {
					if adjusmentEntry.NameAccount == "" {
						continue // Skip jika tidak valid
					}

					// Cek apakah akun cocok
					if generalJournal.NameAccount == adjusmentEntry.NameAccount {
						var kredit, debit float64

						// Penyesuaian debit dan kredit
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

						// Simpan hasil dalam payload
						payload := map[string]any{
							"name_account": generalJournal.NameAccount,
							"kredit":       kredit,
							"debit":        debit,
						}

						worksheetPayloadArray = append(worksheetPayloadArray, payload)
					}
				}
			}

			// Gabungkan worksheet dan data terkait
			payloadResult := map[string]any{
				"worksheet":      worksheet,
				"data_worksheet": worksheetPayloadArray,
			}

			// Tambahkan ke array hasil akhir
			payloadArray = append(payloadArray, payloadResult)
		}

		// Kembalikan hasil dalam bentuk array karena kita mengambil semua worksheet
		return payloadArray, nil
	}

	payload := []map[string]any{
		{
			"worksheet": res,
		},
	}

	// Jika `opt` tidak dipilih, kembalikan hanya data worksheet
	return payload, nil
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
