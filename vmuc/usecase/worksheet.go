package usecase

import (
	"context"
	"time"
	"vmuc-fintech-backend-web-go/domain"
)

type worksheetUseCase struct {
	worksheetRepository domain.WorksheetRepository
	contextTimeout      time.Duration
}

func NewWorksheetUseCase(worksheet domain.WorksheetRepository, t time.Duration) domain.WorksheetUseCase {
	return &worksheetUseCase{
		worksheetRepository: worksheet,
		contextTimeout:      t,
	}
}

func (c *worksheetUseCase) FetchWorksheetByID(ctx context.Context, id uint) (*domain.Worksheet, error) {
	res, err := c.worksheetRepository.RetrieveWorksheetByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
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
