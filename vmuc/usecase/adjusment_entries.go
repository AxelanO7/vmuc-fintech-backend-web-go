package usecase

import (
	"context"
	"time"
	"vmuc-fintech-backend-web-go/domain"
)

type adjusmentEntriesUseCase struct {
	adjusmentEntriesRepository domain.AdjusmentEntriesRepository
	contextTimeout             time.Duration
}

func NewAdjusmentEntriesUseCase(adjusmentEntries domain.AdjusmentEntriesRepository, t time.Duration) domain.AdjusmentEntriesUseCase {
	return &adjusmentEntriesUseCase{
		adjusmentEntriesRepository: adjusmentEntries,
		contextTimeout:             t,
	}
}

func (c *adjusmentEntriesUseCase) FetchAdjusmentEntriesByID(ctx context.Context, id uint) (*domain.AdjusmentEntries, error) {
	res, err := c.adjusmentEntriesRepository.RetrieveAdjusmentEntriesByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *adjusmentEntriesUseCase) FetchAdjusmentEntriess(ctx context.Context) ([]domain.AdjusmentEntries, error) {
	res, err := c.adjusmentEntriesRepository.RetrieveAllAdjusmentEntries()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *adjusmentEntriesUseCase) AddAdjusmentEntries(ctx context.Context, req *domain.AdjusmentEntries) (*domain.AdjusmentEntries, error) {
	res, err := c.adjusmentEntriesRepository.CreateAdjusmentEntries(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *adjusmentEntriesUseCase) AddBulkAdjusmentEntries(ctx context.Context, req []*domain.AdjusmentEntries) ([]*domain.AdjusmentEntries, error) {
	res, err := c.adjusmentEntriesRepository.CreateBulkAdjusmentEntries(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *adjusmentEntriesUseCase) EditAdjusmentEntries(ctx context.Context, req *domain.AdjusmentEntries) (*domain.AdjusmentEntries, error) {
	res, err := c.adjusmentEntriesRepository.UpdateAdjusmentEntries(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *adjusmentEntriesUseCase) EditBulkAdjusmentEntries(ctx context.Context, req []*domain.AdjusmentEntries) ([]*domain.AdjusmentEntries, error) {
	res, err := c.adjusmentEntriesRepository.UpdateBulkAdjusmentEntries(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *adjusmentEntriesUseCase) DeleteAdjusmentEntries(ctx context.Context, id uint) error {
	err := c.adjusmentEntriesRepository.DeleteAdjusmentEntries(id)
	if err != nil {
		return err
	}
	return nil
}
