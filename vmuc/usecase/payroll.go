package usecase

import (
	"context"
	"time"
	"vmuc-fintech-backend-web-go/domain"
)

type payrollUseCase struct {
	payrollRepository domain.PayrollRepository
	contextTimeout    time.Duration
}

func NewPayrollUseCase(payroll domain.PayrollRepository, t time.Duration) domain.PayrollUseCase {
	return &payrollUseCase{
		payrollRepository: payroll,
		contextTimeout:    t,
	}
}

func (c *payrollUseCase) FetchPayrollByID(ctx context.Context, id uint) (*domain.Payroll, error) {
	res, err := c.payrollRepository.RetrievePayrollByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *payrollUseCase) FetchPayrolls(ctx context.Context) ([]domain.Payroll, error) {
	res, err := c.payrollRepository.RetrieveAllPayroll()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *payrollUseCase) AddPayroll(ctx context.Context, req *domain.Payroll) (*domain.Payroll, error) {
	res, err := c.payrollRepository.CreatePayroll(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *payrollUseCase) AddBulkPayroll(ctx context.Context, req []*domain.Payroll) ([]*domain.Payroll, error) {
	res, err := c.payrollRepository.CreateBulkPayroll(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *payrollUseCase) EditPayroll(ctx context.Context, req *domain.Payroll) (*domain.Payroll, error) {
	res, err := c.payrollRepository.UpdatePayroll(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *payrollUseCase) EditBulkPayroll(ctx context.Context, req []*domain.Payroll) ([]*domain.Payroll, error) {
	res, err := c.payrollRepository.UpdateBulkPayroll(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *payrollUseCase) DeletePayroll(ctx context.Context, id uint) error {
	err := c.payrollRepository.DeletePayroll(id)
	if err != nil {
		return err
	}
	return nil
}
