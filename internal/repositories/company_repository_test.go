package repositories

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/mock"
	"github.com/br4tech/auth-nex/internal/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCompanyRepository_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock.NewMockIDatabase[port.IModel](ctrl)
	repo := NewCompanyRepository(mockDB)

	company := &domain.Company{
		Id:                1,
		LegalName:         "Nome Legal LTDA",
		TradeName:         "Nome Fantasia",
		Document:          "12345678901234",
		StateRegistration: "123456789",
		Type:              "LTDA",
		TenantId:          1,
		Schema:            "public",
	}

	companyModel := new(model.Company)
	companyModel.FromDomain(company)

	t.Run("Success", func(t *testing.T) {
		mockDB.EXPECT().Create(companyModel).Return(company.Id, nil)

		result, err := repo.Create(company)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("Error", func(t *testing.T) {
		expectedError := errors.New("database error")

		mockDB.EXPECT().Create(gomock.Any()).Return(0, expectedError)

		_, err := repo.Create(company)
		assert.Error(t, err)
		assert.EqualError(t, err, "database error")
	})
}
