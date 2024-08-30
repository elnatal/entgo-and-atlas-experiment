package service

import (
	"context"
	"errors"

	"github.com/elnatal/go-experiment/internal/core/domain"
	"github.com/elnatal/go-experiment/internal/core/port"
	"github.com/elnatal/go-experiment/internal/core/util"
)

type UserService struct {
	repo port.UserRepository
}

// NewUserService creates a new user service instance
func NewUserService(repo port.UserRepository) *UserService {
	return &UserService{
		repo,
	}
}

// Register creates a new user
func (us *UserService) Register(ctx context.Context, user *domain.User) (*domain.User, error) {
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword

	user, err = us.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUser gets a user by ID
func (us *UserService) GetUser(ctx context.Context, id int) (*domain.User, error) {
	user, err := us.repo.GetUserByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// ListUsers lists all users
func (us *UserService) ListUsers(ctx context.Context, skip, limit int) ([]domain.User, error) {
	users, err := us.repo.ListUsers(ctx, skip, limit)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// UpdateUser updates a user's name, email, and password
func (us *UserService) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	existingUser, err := us.repo.GetUserByID(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	emptyData := user.Name == "" &&
		user.Email == "" &&
		user.Password == ""

	sameData := existingUser.Name == user.Name &&
		existingUser.Email == user.Email
	if emptyData || sameData {
		return nil, errors.New("nothing to update")
	}

	var hashedPassword string

	if user.Password != "" {
		hashedPassword, err = util.HashPassword(user.Password)
		if err != nil {
			return nil, err
		}
	}

	user.Password = hashedPassword

	_, err = us.repo.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser deletes a user by ID
func (us *UserService) DeleteUser(ctx context.Context, id int) error {
	_, err := us.repo.GetUserByID(ctx, id)
	if err != nil {
		return err
	}

	return us.repo.DeleteUser(ctx, id)
}
