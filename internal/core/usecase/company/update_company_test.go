package company

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/test/factories"
	"github.com/br4tech/auth-nex/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUpdateCompanyUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	companyRepoMock := mock.NewMockICompanyRepository(ctrl)
	updateCompanyUseCase := NewUpdateCompanyUseCase(companyRepoMock)

	companyToUpdate := factories.NewCompanyFactory() // Cria uma empresa de teste com a fábrica
	updatedCompany := factories.NewCompanyFactory()  // Cria uma empresa atualizada com a fábrica

	t.Run("Success", func(t *testing.T) {
		companyRepoMock.EXPECT().Update(companyToUpdate).Return(updatedCompany, nil)

		company, err := updateCompanyUseCase.Execute(companyToUpdate)

		assert.NoError(t, err)
		assert.NotNil(t, company)
		assert.Equal(t, updatedCompany, company)
	})

	t.Run("Error", func(t *testing.T) {
		expectedError := errors.New("Erro ao atualizar empresa")

		companyRepoMock.EXPECT().Update(companyToUpdate).Return(nil, expectedError)

		company, err := updateCompanyUseCase.Execute(companyToUpdate)

		assert.Error(t, err)
		assert.Nil(t, company)
		assert.Equal(t, expectedError, err)
	})
}
