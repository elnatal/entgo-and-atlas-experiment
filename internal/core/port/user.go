package port

import (
	"context"

	"github.com/elnatal/go-experiment/internal/core/domain"
)

//go:generate mockgen -source=user.go -destination=mock/user.go -package=mock

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUserByID(ctx context.Context, id int) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	ListUsers(ctx context.Context, skip, limit int) ([]domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type UserService interface {
	Register(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUser(ctx context.Context, id int) (*domain.User, error)
	ListUsers(ctx context.Context, skip, limit int) ([]domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, id int) error
}
