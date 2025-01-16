package usecase

import (
	"context"
	"time"
	"vmuc-fintech-backend-web-go/domain"
)

type trialBalanceUseCase struct {
	trialBalanceRepository domain.TrialBalanceRepository
	contextTimeout         time.Duration
}

func NewTrialBalanceUseCase(trialBalance domain.TrialBalanceRepository, t time.Duration) domain.TrialBalanceUseCase {
	return &trialBalanceUseCase{
		trialBalanceRepository: trialBalance,
		contextTimeout:         t,
	}
}

func (c *trialBalanceUseCase) FetchTrialBalanceByID(ctx context.Context, id uint) (*domain.TrialBalance, error) {
	res, err := c.trialBalanceRepository.RetrieveTrialBalanceByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *trialBalanceUseCase) FetchTrialBalances(ctx context.Context) ([]domain.TrialBalance, error) {
	res, err := c.trialBalanceRepository.RetrieveAllTrialBalance()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *trialBalanceUseCase) AddTrialBalance(ctx context.Context, req *domain.TrialBalance) (*domain.TrialBalance, error) {
	res, err := c.trialBalanceRepository.CreateTrialBalance(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *trialBalanceUseCase) AddBulkTrialBalance(ctx context.Context, req []*domain.TrialBalance) ([]*domain.TrialBalance, error) {
	res, err := c.trialBalanceRepository.CreateBulkTrialBalance(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *trialBalanceUseCase) EditTrialBalance(ctx context.Context, req *domain.TrialBalance) (*domain.TrialBalance, error) {
	res, err := c.trialBalanceRepository.UpdateTrialBalance(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *trialBalanceUseCase) EditBulkTrialBalance(ctx context.Context, req []*domain.TrialBalance) ([]*domain.TrialBalance, error) {
	res, err := c.trialBalanceRepository.UpdateBulkTrialBalance(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *trialBalanceUseCase) DeleteTrialBalance(ctx context.Context, id uint) error {
	err := c.trialBalanceRepository.DeleteTrialBalance(id)
	if err != nil {
		return err
	}
	return nil
}
