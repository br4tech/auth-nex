package customer

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/test/factories"
	"github.com/br4tech/auth-nex/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateUserCustomerUsecase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mock.NewMockIUserRepository(ctrl)
	createUserCustomerUseCase := NewCreateUserCustomerUsecase(userRepoMock)

	userToCreate := factories.NewUserFactory("customer")
	createdUser := factories.NewUserFactory("customer")

	t.Run("Success", func(t *testing.T) {
		createdUser.Role = "customer"

		userRepoMock.EXPECT().Create(gomock.Any()).Return(createdUser, nil)

		err := createUserCustomerUseCase.Execute(userToCreate)

		assert.NoError(t, err)
		assert.Equal(t, "customer", userToCreate.Role)
	})

	t.Run("Error", func(t *testing.T) {
		expectedError := errors.New("Erro ao criar usuário customer")

		userRepoMock.EXPECT().Create(userToCreate).Return(nil, expectedError)

		err := createUserCustomerUseCase.Execute(userToCreate)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}
