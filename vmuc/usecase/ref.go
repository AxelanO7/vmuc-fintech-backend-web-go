package usecase

import (
	"context"
	"time"
	"vmuc-fintech-backend-web-go/domain"
)

type refUseCase struct {
	refRepository  domain.RefRepository
	contextTimeout time.Duration
}

func NewRefUseCase(ref domain.RefRepository, t time.Duration) domain.RefUseCase {
	return &refUseCase{
		refRepository:  ref,
		contextTimeout: t,
	}
}

func (c *refUseCase) FetchRefByID(ctx context.Context, id uint) (*domain.Ref, error) {
	res, err := c.refRepository.RetrieveRefByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *refUseCase) FetchRefs(ctx context.Context) ([]domain.Ref, error) {
	res, err := c.refRepository.RetrieveRefs()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *refUseCase) AddRef(ctx context.Context, req *domain.Ref) (*domain.Ref, error) {
	res, err := c.refRepository.CreateRef(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *refUseCase) EditRef(ctx context.Context, req *domain.Ref) (*domain.Ref, error) {
	res, err := c.refRepository.UpdateRef(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *refUseCase) DeleteRef(ctx context.Context, id uint) error {
	err := c.refRepository.DeleteRef(id)
	if err != nil {
		return err
	}
	return nil
}
