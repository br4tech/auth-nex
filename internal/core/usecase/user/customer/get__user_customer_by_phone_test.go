package customer

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/test/factories"
	"github.com/br4tech/auth-nex/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetUserCustomerByPhoneUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mock.NewMockIUserRepository(ctrl)
	getUserCustomerByPhoneUseCase := NewGetUserCustomerByPhoneUseCase(userRepoMock)

	expectedUser := factories.NewUserFactory("customer")

	t.Run("Success", func(t *testing.T) {
		userPhone := "+55 11 987654321"
		expectedUser.Phone = userPhone

		filter := map[string]interface{}{"phone": userPhone}

		userRepoMock.EXPECT().FindBy(filter).Return(expectedUser, nil)

		user, err := getUserCustomerByPhoneUseCase.Execute(userPhone)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, expectedUser, user)
	})

	t.Run("Error - User not found", func(t *testing.T) {
		userPhone := "+55 11 123456789"
		expectedError := errors.New("Usuário não encontrado")

		filter := map[string]interface{}{"phone": userPhone}

		userRepoMock.EXPECT().FindBy(filter).Return(nil, expectedError)

		user, err := getUserCustomerByPhoneUseCase.Execute(userPhone)

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, expectedError, err)
	})
}
