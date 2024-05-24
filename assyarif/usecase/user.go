package usecase

import (
	"assyarif-backend-web-go/domain"
	"assyarif-backend-web-go/middleware"
	"assyarif-backend-web-go/utils"
	"context"
	"fmt"
	"time"
)

type userUseCase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUseCase(usr domain.UserRepository, t time.Duration) domain.UserUseCase {
	return &userUseCase{
		userRepository: usr,
		contextTimeout: t,
	}
}

func (c *userUseCase) LoginUser(ctx context.Context, req *domain.LoginPayload) (*domain.User, string, error) {
	res, err := c.userRepository.GetUser(*req.Username)
	if err != nil {
		return nil, "", err
	}
	err = utils.VerifyPassword(req.Password, res.Password)
	if err != nil {
		return nil, "", fmt.Errorf("error verifying password: %v", err)
	}
	tokPay := domain.TokenClaims{
		User: res,
	}
	token, err := middleware.CreateToken(&tokPay)
	if err != nil {
		return nil, "", fmt.Errorf("cannot create token: %v", err)
	}
	return res, token, nil
}
