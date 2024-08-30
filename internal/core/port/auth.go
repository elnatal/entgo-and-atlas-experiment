package port

import (
	"context"

	"github.com/elnatal/go-experiment/internal/core/domain"
)

//go:generate mockgen -source=auth.go -destination=mock/auth.go -package=mock

type TokenService interface {
	CreateToken(user *domain.User) (string, error)
	VerifyToken(token string) (*domain.TokenPayload, error)
}

type AuthService interface {
	Login(ctx context.Context, email, password string) (string, error)
}
