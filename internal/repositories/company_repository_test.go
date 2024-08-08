package repositories

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/mock"
	"github.com/br4tech/auth-nex/internal/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCompanyRepository_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock.NewMockIDatabase[model.Company](ctrl)
	repo := NewCompanyRepository(mockDB)

	company := &domain.Company{
		LegalName:         "Nome Legal LTDA",
		TradeName:         "Nome Fantasia",
		Document:          "12345678901234",
		StateRegistration: "123456789",
		Address: domain.Address{
			Street:     "Rua Exemplo",
			Number:     "123",
			Complement: "Apto 456",
			District:   "Bairro Legal",
			City:       "Cidade Grande",
			State:      "SP",
			ZipCode:    "12345-678",
		},
		Type:     "LTDA",
		TenantId: 1,
		Schema:   "public",
		Partners: []domain.Partner{
			{
				CompanyId:     1,
				Participation: 50,
				UserId:        1,
			},
		},
	}

	companyModel := new(model.Company)
	companyModel.FromDomain(company)

	t.Run("Success", func(t *testing.T) {
		mockDB.EXPECT().Create(companyModel).Return(company.Id, nil)

		result, err := repo.Create(company)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, company.Id, result.Id)
	})

	t.Run("Error", func(t *testing.T) {
		expectedError := errors.New("database error")

		mockDB.EXPECT().Create(gomock.Any()).Return(nil, expectedError)

		_, err := repo.Create(company)

		if err == nil {
			t.Fatal("CreateCompany did not return an error")
		}

		if err != expectedError {
			t.Fatalf("CreateCompany returned unexpected error: %v", err)
		}
	})
}
