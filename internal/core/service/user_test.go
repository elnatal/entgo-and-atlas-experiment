package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/elnatal/go-experiment/internal/core/domain"
	"github.com/elnatal/go-experiment/internal/core/port/mock"
	"github.com/elnatal/go-experiment/internal/core/service"
	"github.com/elnatal/go-experiment/internal/core/util"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type registerTestedInput struct {
	user *domain.User
}

type registerExpectedOutput struct {
	user *domain.User
	err  error
}

func TestUserService_Register(t *testing.T) {
	ctx := context.Background()
	userName := gofakeit.Name()
	userEmail := gofakeit.Email()
	userPassword := gofakeit.Password(true, true, true, true, false, 8)
	hashedPassword, _ := util.HashPassword(userPassword)

	userInput := &domain.User{
		Name:     userName,
		Email:    userEmail,
		Password: userPassword,
	}
	userOutput := &domain.User{
		ID:       int(gofakeit.Int64()),
		Name:     userName,
		Email:    userEmail,
		Password: hashedPassword,
	}

	testCases := []struct {
		desc  string
		mocks func(
			userRepo *mock.MockUserRepository,
		)
		input    registerTestedInput
		expected registerExpectedOutput
	}{
		{
			desc: "Success",
			mocks: func(
				userRepo *mock.MockUserRepository,
			) {
				userRepo.EXPECT().
					CreateUser(gomock.Any(), gomock.Eq(userInput)).
					Return(userOutput, nil)
			},
			input: registerTestedInput{
				user: userInput,
			},
			expected: registerExpectedOutput{
				user: userOutput,
				err:  nil,
			},
		},
		{
			desc: "Fail_DuplicateData",
			mocks: func(
				userRepo *mock.MockUserRepository,
			) {
				userRepo.EXPECT().
					CreateUser(gomock.Any(), gomock.Eq(userInput)).
					Return(nil, errors.New("user already exist"))
			},
			input: registerTestedInput{
				user: userInput,
			},
			expected: registerExpectedOutput{
				user: nil,
				err:  errors.New("user already exist"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userRepo := mock.NewMockUserRepository(ctrl)

			tc.mocks(userRepo)

			userService := service.NewUserService(userRepo)

			user, err := userService.Register(ctx, tc.input.user)
			assert.Equal(t, tc.expected.err, err, "Error mismatch")
			assert.Equal(t, tc.expected.user, user, "User mismatch")
		})
	}
}
