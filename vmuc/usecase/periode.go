package usecase

import (
	"context"
	"time"
	"vmuc-fintech-backend-web-go/domain"
)

type periodeUseCase struct {
	periodeRepository domain.PeriodeRepository
	payrollRepository domain.PayrollRepository
	contextTimeout    time.Duration
}

func NewPeriodeUseCase(payroll domain.PeriodeRepository, payrep domain.PayrollRepository, t time.Duration) domain.PeriodeUseCase {
	return &periodeUseCase{
		periodeRepository: payroll,
		payrollRepository: payrep,
		contextTimeout:    t,
	}
}

func (c *periodeUseCase) FetchPeriodeByID(ctx context.Context, id uint) (*domain.Periode, error) {
	return c.periodeRepository.RetrievePeriodeByID(id)
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

func (c *periodeUseCase) AddPeriode(ctx context.Context, req *domain.Periode) (*domain.Periode, error) {
	return c.periodeRepository.CreatePeriode(req)
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
