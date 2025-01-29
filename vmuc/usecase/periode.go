package usecase

import (
	"context"
	"fmt"
	"time"
	"vmuc-fintech-backend-web-go/domain"
)

type periodeUseCase struct {
	periodeRepository          domain.PeriodeRepository
	payrollRepository          domain.PayrollRepository
	adjusmentEntriesRepository domain.AdjusmentEntriesRepository
	generalJournalRepository   domain.GeneralJournalRepository
	trialBalanceRepository     domain.TrialBalanceRepository
	contextTimeout             time.Duration
}

func NewPeriodeUseCase(payroll domain.PeriodeRepository, payrep domain.PayrollRepository, adrep domain.AdjusmentEntriesRepository, genrep domain.GeneralJournalRepository, trirep domain.TrialBalanceRepository, t time.Duration) domain.PeriodeUseCase {
	return &periodeUseCase{
		periodeRepository:          payroll,
		payrollRepository:          payrep,
		adjusmentEntriesRepository: adrep,
		generalJournalRepository:   genrep,
		trialBalanceRepository:     trirep,
		contextTimeout:             t,
	}
}

func (c *periodeUseCase) FetchPeriodeByID(ctx context.Context, id uint) (*domain.Periode, error) {
	return c.periodeRepository.RetrievePeriodeByID(id)
}

func (c *periodeUseCase) GetTrialBalanceReportByPeriode(ctx context.Context, periode string) (map[string]any, error) {
	res, err := c.periodeRepository.GetPeriodeByPeriode(periode)
	if err != nil {
		return nil, err
	}

	trialBalance, err := c.generalJournalRepository.GetGeneralJournalByGeneralJournalPeriodeId(res.ID)
	if err != nil {
		return nil, err
	}

	payload := map[string]any{
		"period":        res,
		"trial_balance": trialBalance,
	}
	return payload, nil
}

func (c *periodeUseCase) FetchPayrollPeriode(ctx context.Context) ([]domain.Periode, error) {
	// Ambil semua data PayrollPeriode
	payrollPeriodes, err := c.periodeRepository.RetrieveAllPeriode()
	if err != nil {
		return nil, err
	}

	// Ambil data Payroll untuk setiap PayrollPeriode
	for i := range payrollPeriodes {
		payrolls, err := c.payrollRepository.GetPayrollByPayrollPeriodeId(payrollPeriodes[i].ID)
		if err != nil {
			return nil, err
		}
		payrollPeriodes[i].Payrolls = payrolls
	}

	return payrollPeriodes, nil
}

func (c *periodeUseCase) FetchAdjusmentEntriesPeriode(ctx context.Context) ([]domain.Periode, error) {
	// Ambil semua data PayrollPeriode
	payrollPeriodes, err := c.periodeRepository.RetrieveAllPeriode()
	if err != nil {
		return nil, err
	}

	// Ambil data Payroll untuk setiap PayrollPeriode
	for i := range payrollPeriodes {
		payrolls, err := c.adjusmentEntriesRepository.GetAdjusmentEntriesByAdjusmentEntriesPeriodeId(payrollPeriodes[i].ID)
		if err != nil {
			return nil, err
		}
		payrollPeriodes[i].AdjusmentEntries = payrolls
	}

	return payrollPeriodes, nil
}

func (c *periodeUseCase) FetchGeneralJournalPeriode(ctx context.Context) ([]domain.Periode, error) {
	// Ambil semua data PayrollPeriode
	payrollPeriodes, err := c.periodeRepository.RetrieveAllPeriode()
	if err != nil {
		return nil, err
	}

	// Ambil data Payroll untuk setiap PayrollPeriode
	for i := range payrollPeriodes {
		payrolls, err := c.generalJournalRepository.GetGeneralJournalByGeneralJournalPeriodeId(payrollPeriodes[i].ID)
		if err != nil {
			return nil, err
		}
		payrollPeriodes[i].GeneralJournal = payrolls
	}

	return payrollPeriodes, nil
}

func (c *periodeUseCase) FetchTrialBalancePeriode(ctx context.Context) ([]domain.Periode, error) {
	// Ambil semua data PayrollPeriode
	payrollPeriodes, err := c.periodeRepository.RetrieveAllPeriode()
	if err != nil {
		return nil, err
	}

	// Ambil data Payroll untuk setiap PayrollPeriode
	for i := range payrollPeriodes {
		payrolls, err := c.trialBalanceRepository.GetTrialBalanceByTrialBalancePeriodeId(payrollPeriodes[i].ID)
		if err != nil {
			return nil, err
		}
		payrollPeriodes[i].TrialBalance = payrolls
	}

	return payrollPeriodes, nil
}

func (c *periodeUseCase) AddPeriode(ctx context.Context, req *domain.Periode) (*domain.Periode, error) {
	fmt.Println("request", req)
	// return c.periodeRepository.CreatePeriode(req)
	periodes, err := c.periodeRepository.RetrieveAllPeriode()
	if err != nil {
		return nil, err
	}
	fmt.Println("periodes", periodes)

	var resPeriode *domain.Periode

	// check if period is empty
	if len(periodes) == 0 {
		resPeriode, err = c.periodeRepository.CreatePeriode(&domain.Periode{
			Period:      req.Period,
			Description: req.Description,
		})
		if err != nil {
			return nil, err
		}
	}
	// check if period already exists
	for i := range periodes {
		// if periode already exists, return the periode
		if periodes[i].Period == req.Period {
			resPeriode = &periodes[i]
			break
		}
		// if periode does not exist, create new periode
		if i == len(periodes)-1 {
			resPeriode, err = c.periodeRepository.CreatePeriode(&domain.Periode{
				Period:      req.Period,
				Description: req.Description,
			})
			if err != nil {
				return nil, err
			}
		}
	}
	fmt.Println("resPeriode", resPeriode)

	if req.Payrolls != nil && len(req.Payrolls) > 0 {
		fmt.Println("payroll", req.Payrolls)
		for _, val := range req.Payrolls {
			val.IdPeriode = resPeriode.ID
			fmt.Println("payroll for", val)
			_, err := c.payrollRepository.CreatePayroll(&val)
			if err != nil {
				return nil, err
			}
			resPeriode.Payrolls = append(resPeriode.Payrolls, val)
		}
	}
	if req.AdjusmentEntries != nil && len(req.AdjusmentEntries) > 0 {
		fmt.Println("adjusment", req.AdjusmentEntries)
		for _, val := range req.AdjusmentEntries {
			val.IdPeriode = resPeriode.ID
			fmt.Println("adjusment for", val)
			_, err := c.adjusmentEntriesRepository.CreateAdjusmentEntries(&val)
			if err != nil {
				return nil, err
			}
			resPeriode.AdjusmentEntries = append(resPeriode.AdjusmentEntries, val)
		}
	}
	if req.GeneralJournal != nil && len(req.GeneralJournal) > 0 {
		fmt.Println("general", req.GeneralJournal)
		for _, val := range req.GeneralJournal {
			val.IdPeriode = resPeriode.ID
			fmt.Println("general for", val)
			_, err := c.generalJournalRepository.CreateGeneralJournal(&val)
			if err != nil {
				return nil, err
			}
			resPeriode.GeneralJournal = append(resPeriode.GeneralJournal, val)
		}
	}
	if req.TrialBalance != nil && len(req.TrialBalance) > 0 {
		fmt.Println("trial", req.TrialBalance)
		for _, val := range req.TrialBalance {
			val.IdPeriode = resPeriode.ID
			fmt.Println("trial for", val)
			_, err := c.trialBalanceRepository.CreateTrialBalance(&val)
			if err != nil {
				return nil, err
			}
			resPeriode.TrialBalance = append(resPeriode.TrialBalance, val)
		}
	}
	return resPeriode, nil
}

func (c *periodeUseCase) AddBulkPeriode(ctx context.Context, req []*domain.Periode) ([]*domain.Periode, error) {
	return c.periodeRepository.CreateBulkPeriode(req)
}

func (c *periodeUseCase) EditPeriode(ctx context.Context, req *domain.Periode) (*domain.Periode, error) {
	res, err := c.periodeRepository.UpdatePeriode(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *periodeUseCase) EditBulkPeriode(ctx context.Context, req []*domain.Periode) ([]*domain.Periode, error) {
	res, err := c.periodeRepository.UpdateBulkPeriode(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *periodeUseCase) DeletePeriode(ctx context.Context, id uint) error {
	err := c.periodeRepository.DeletePeriode(id)
	if err != nil {
		return err
	}
	return nil
}

func (c *periodeUseCase) AddPayrollPeriode(ctx context.Context, req *domain.PayrollPeriode) (*domain.PayrollPeriode, error) {
	// periods, err := c.periodeRepository.RetrieveAllPeriode()
	// if err != nil {
	// 	return nil, fmt.Errorf("error retrieving all periods: %w", err)
	// }

	// var period *domain.Periode
	// // check if period already exists
	// for i := range periods {
	// 	if periods[i].Period == req.Period {
	// 		period = &periods[i]
	// 		break
	// 	}
	// }

	// // if period does not exist, create new period
	// if period == nil {
	// 	period, err = c.periodeRepository.CreatePeriode(&domain.Periode{
	// 		Period:      req.Period,
	// 		Description: req.Description,
	// 	})
	// 	if err != nil {
	// 		return nil, fmt.Errorf("error creating period: %w", err)
	// 	}
	// }

	// // change period id to new period id
	// for i := range req.Payrolls {
	// 	req.Payrolls[i].IdPeriode = period.ID
	// }

	// // check if payrolls is empty
	// if len(req.Payrolls) == 0 {
	// 	return req, nil
	// }

	// // create bulk payroll
	// payrolls := make([]*domain.Payroll, len(req.Payrolls))
	// for i := range req.Payrolls {
	// 	payrolls[i] = &req.Payrolls[i]
	// }
	// _, err = c.payrollRepository.CreateBulkPayroll(payrolls)
	// if err != nil {
	// 	return nil, fmt.Errorf("error creating bulk payroll: %w", err)
	// }
	// return req, nil

	var idPeriode uint

	res, err := c.periodeRepository.GetPeriodeByPeriode(req.Period)
	idPeriode = res.ID
	if err != nil {
		res, err := c.periodeRepository.CreatePeriode(&domain.Periode{
			Period:      req.Period,
			Description: req.Description,
		})
		if err != nil {
			return nil, fmt.Errorf("error creating period: %w", err)
		}
		idPeriode = res.ID
	}

	for _, val := range req.Payrolls {
		val.IdPeriode = idPeriode
		_, err := c.payrollRepository.CreatePayroll(&val)
		if err != nil {
			return nil, err
		}

		payload := domain.GeneralJournal{
			IdPeriode:   idPeriode,
			NameAccount: "beban gaji",
			Date:        time.Now().Format(time.RFC3339),
			Information: "beban gaji",
			Kredit:      float64(val.Total),
			IdRef:       int(req.IdRef),
		}

		_, err = c.generalJournalRepository.CreateGeneralJournal(&payload)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}
