package usecase

import (
	"context"
	"time"
	"vmuc-fintech-backend-web-go/domain"
)

type payrollPeriodeUseCase struct {
	payrollPeriodeRepository domain.PayrollPeriodeRepository
	payrollRepository        domain.PayrollRepository
	contextTimeout           time.Duration
}

func NewPayrollPeriodeUseCase(payroll domain.PayrollPeriodeRepository, payrep domain.PayrollRepository, t time.Duration) domain.PayrollPeriodeUseCase {
	return &payrollPeriodeUseCase{
		payrollPeriodeRepository: payroll,
		payrollRepository:        payrep,
		contextTimeout:           t,
	}
}

func (c *payrollPeriodeUseCase) FetchPayrollPeriodeByID(ctx context.Context, id uint) (*domain.PayrollPeriode, error) {
	return c.payrollPeriodeRepository.RetrievePayrollPeriodeByID(id)
}

func (c *payrollPeriodeUseCase) FetchPayrollPeriode(ctx context.Context) ([]domain.PayrollPeriode, error) {
	// Ambil semua data PayrollPeriode
	payrollPeriodes, err := c.payrollPeriodeRepository.RetrieveAllPayrollPeriode()
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

func (c *payrollPeriodeUseCase) AddPayrollPeriode(ctx context.Context, req *domain.PayrollPeriode) (*domain.PayrollPeriode, error) {
	res, err := c.payrollPeriodeRepository.CreatePayrollPeriode(req)
	if err != nil {
		return nil, err
	}
	payrolls := make([]*domain.Payroll, 0)
	for i := range req.Payrolls {
		req.Payrolls[i].IdPayrollPeriode = res.ID
		payrolls = append(payrolls, &req.Payrolls[i])
	}
	_, err = c.payrollRepository.CreateBulkPayroll(payrolls)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *payrollPeriodeUseCase) AddBulkPayrollPeriode(ctx context.Context, req []*domain.PayrollPeriode) ([]*domain.PayrollPeriode, error) {
	res, err := c.payrollPeriodeRepository.CreateBulkPayrollPeriode(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *payrollPeriodeUseCase) EditPayrollPeriode(ctx context.Context, req *domain.PayrollPeriode) (*domain.PayrollPeriode, error) {
	res, err := c.payrollPeriodeRepository.UpdatePayrollPeriode(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *payrollPeriodeUseCase) EditBulkPayrollPeriode(ctx context.Context, req []*domain.PayrollPeriode) ([]*domain.PayrollPeriode, error) {
	res, err := c.payrollPeriodeRepository.UpdateBulkPayrollPeriode(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *payrollPeriodeUseCase) DeletePayrollPeriode(ctx context.Context, id uint) error {
	err := c.payrollPeriodeRepository.DeletePayrollPeriode(id)
	if err != nil {
		return err
	}
	return nil
}
