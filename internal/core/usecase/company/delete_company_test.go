package company

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/test/mock" // Certifique-se de que o caminho está correto
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestDeleteCompanyUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	companyRepoMock := mock.NewMockICompanyRepository(ctrl)
	deleteCompanyUseCase := NewDeleteCompanyUseCase(companyRepoMock)

	t.Run("Success", func(t *testing.T) {
		companyId := 1

		companyRepoMock.EXPECT().Delete(companyId).Return(nil)

		err := deleteCompanyUseCase.Execute(companyId)

		assert.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		companyId := 1
		expectedError := errors.New("Erro ao excluir empresa")

		companyRepoMock.EXPECT().Delete(companyId).Return(expectedError)

		err := deleteCompanyUseCase.Execute(companyId)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}
