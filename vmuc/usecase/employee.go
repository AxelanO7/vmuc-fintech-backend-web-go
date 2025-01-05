package usecase

import (
	"context"
	"fmt"
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

func (c *employeeUseCase) CreateEmployee(ctx context.Context, req *domain.Employee) (*domain.Employee, error) {
	res, err := c.employeeRepository.CreateEmployee(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *employeeUseCase) UpdateEmployee(ctx context.Context, req *domain.Employee) (*domain.Employee, error) {
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

func (c *employeeUseCase) ShowEmployeeLastNumber(ctx context.Context) (int, error) {
	var res []domain.Employee
	res, err := c.employeeRepository.RetrieveAllEmployee()
	if err != nil {
		return 0, err
	}
	lastNumber := 0
	for _, v := range res {
		if v.ID > uint(lastNumber) {
			lastNumber = int(v.ID)
		}
	}
	fmt.Println(lastNumber)
	return lastNumber, nil
}
