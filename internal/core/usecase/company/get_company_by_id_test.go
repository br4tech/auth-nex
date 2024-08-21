package company

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/test/factories"
	"github.com/br4tech/auth-nex/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetCompanyByIdUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	companyRepoMock := mock.NewMockICompanyRepository(ctrl)
	getCompanyByIdUseCase := NewGetCompanyByIdUsecase(companyRepoMock)

	t.Run("Success", func(t *testing.T) {
		companyId := 1
		expectedCompany := factories.NewCompanyFactory()

		companyRepoMock.EXPECT().FindById(companyId).Return(expectedCompany, nil)

		company, err := getCompanyByIdUseCase.Execute(companyId)

		assert.NoError(t, err)
		assert.NotNil(t, company)
		assert.Equal(t, expectedCompany, company)
	})

	t.Run("Error - Company not found", func(t *testing.T) {
		companyId := 1
		expectedError := errors.New("Empresa não encontrada")

		companyRepoMock.EXPECT().FindById(companyId).Return(nil, expectedError)

		company, err := getCompanyByIdUseCase.Execute(companyId)

		assert.Error(t, err)
		assert.Nil(t, company)
		assert.Equal(t, expectedError, err)
	})
}
