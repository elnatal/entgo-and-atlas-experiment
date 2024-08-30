package repository

import (
	"context"

	"github.com/elnatal/go-experiment/internal/adopter/storage/ent"
	"github.com/elnatal/go-experiment/internal/core/domain"

	entUser "github.com/elnatal/go-experiment/ent/user"
)

type UserRepository struct {
	ent *ent.Ent
}

// NewUserRepository creates a new user repository instance
func NewUserRepository(ent *ent.Ent) *UserRepository {
	return &UserRepository{
		ent,
	}
}

// CreateUser creates a new user in the database
func (ur *UserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	u, err := ur.ent.Client.User.
		Create().
		SetName(user.Name).
		SetEmail(user.Email).
		SetPassword(user.Password).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	user.ID = u.ID

	return user, nil
}

func (ur *UserRepository) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	u, err := ur.ent.Client.User.
		Get(ctx, id)

	if err != nil {
		return nil, err
	}

	user := domain.User{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}

	return &user, nil
}

func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	u, err := ur.ent.Client.User.
		Query().
		Where(entUser.Email(email)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	user := domain.User{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}

	return &user, nil
}

func (ur *UserRepository) ListUsers(ctx context.Context, skip, limit int) ([]domain.User, error) {
	u, err := ur.ent.Client.User.
		Query().
		Order(entUser.ByID()).
		Offset(int(skip)).
		Limit(int(limit)).
		All(ctx)

	if err != nil {
		return nil, err
	}

	var users []domain.User

	for _, user := range u {
		users = append(users, domain.User{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		})
	}

	return users, nil
}

func (ur *UserRepository) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	u, err := ur.ent.Client.User.
		UpdateOneID(user.ID).
		SetName(user.Name).
		SetEmail(user.Email).
		SetPassword(user.Password).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	newUser := domain.User{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}

	return &newUser, nil
}

func (ur *UserRepository) DeleteUser(ctx context.Context, id int) error {
	err := ur.ent.Client.User.
		DeleteOneID(id).
		Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}
