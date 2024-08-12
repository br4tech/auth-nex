package company

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCompanyUseCase_CreateCompany(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	companyRepoMock := mock.NewMockICompanyRepository(ctrl)
	companyUsecase := NewCompanyUseCase(companyRepoMock)

	companyDomain := &domain.Company{
		LegalName:         "Empresa Legal Ltda.",
		TradeName:         "Empresa Legal",
		Document:          "12345678901234",
		StateRegistration: "123456789",
		Address:           domain.Address{},
		Partners:          []domain.Partner{},
		Activities:        []domain.Activity{},
		Type:              "LTDA",
		TenantId:          1,
		Schema:            "app-0000",
		Users:             []domain.User{},
	}

	t.Run("success", func(t *testing.T) {
		companyRepoMock.EXPECT().Create(gomock.Any()).Return(companyDomain, nil)

		createdCompany, _ := companyUsecase.Create(companyDomain)

		assert.Equal(t, companyDomain.LegalName, createdCompany.LegalName)
	})

	t.Run("error", func(t *testing.T) {
		companyInvalid := errors.New("Error in company")
		companyRepoMock.EXPECT().Create(gomock.Any()).Return(nil, companyInvalid)

		createdCompany, err := companyUsecase.Create(companyDomain)

		assert.Error(t, err)
		assert.Nil(t, createdCompany)
	})

}
