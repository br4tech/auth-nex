package company

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/test/factories"
	"github.com/br4tech/auth-nex/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateCompanyUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	companyRepoMock := mock.NewMockICompanyRepository(ctrl)
	createCompanyUseCase := NewCreateCompanyUseCase(companyRepoMock)
	expectedCompany := factories.NewCompanyFactory()

	t.Run("Sucess", func(t *testing.T) {
		companyRepoMock.EXPECT().Create(gomock.Any()).Return(expectedCompany, nil)

		company, _ := createCompanyUseCase.Execute(expectedCompany)

		assert.NotNil(t, company)
		assert.Equal(t, expectedCompany, company)
	})

	t.Run("eror", func(t *testing.T) {
		expectedError := errors.New("Error in company")

		companyRepoMock.EXPECT().Create(gomock.Any()).Return(nil, expectedError)

		company, err := createCompanyUseCase.Execute(expectedCompany)

		assert.Error(t, err)
		assert.Nil(t, company)
		assert.Equal(t, expectedError, err)
	})
}
