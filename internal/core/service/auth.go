package service

import (
	"context"

	"github.com/elnatal/go-experiment/internal/core/port"
	"github.com/elnatal/go-experiment/internal/core/util"
)

type AuthService struct {
	repo port.UserRepository
	ts   port.TokenService
}

func NewAuthService(repo port.UserRepository, ts port.TokenService) *AuthService {
	return &AuthService{
		repo,
		ts,
	}
}

func (as *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := as.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	err = util.ComparePassword(password, user.Password)
	if err != nil {
		return "", err
	}

	accessToken, err := as.ts.CreateToken(user)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
