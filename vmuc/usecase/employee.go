package usecase

import (
	"context"
	"time"
	"vmuc-fintech-backend-web-go/domain"
)

type employeeUseCase struct {
	employeeRepository domain.EmployeeRepository
	contextTimeout     time.Duration
}

func NewEmployeeUseCase(employee domain.EmployeeRepository, t time.Duration) domain.EmployeeUseCase {
	return &employeeUseCase{
		employeeRepository: employee,
		contextTimeout:     t,
	}
}

func (c *employeeUseCase) FetchEmployeeByID(ctx context.Context, id uint) (*domain.Employee, error) {
	res, err := c.employeeRepository.RetrieveEmployeeByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *employeeUseCase) FetchEmployees(ctx context.Context) ([]domain.Employee, error) {
	res, err := c.employeeRepository.RetrieveAllEmployee()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *employeeUseCase) AddEmployee(ctx context.Context, req *domain.Employee) (*domain.Employee, error) {
	res, err := c.employeeRepository.CreateEmployee(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *employeeUseCase) EditEmployee(ctx context.Context, req *domain.Employee) (*domain.Employee, error) {
	res, err := c.employeeRepository.UpdateEmployee(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *employeeUseCase) DeleteEmployee(ctx context.Context, id uint) error {
	err := c.employeeRepository.DeleteEmployee(id)
	if err != nil {
		return err
	}
	return nil
}
