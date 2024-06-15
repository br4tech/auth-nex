package repository

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/mock"
	"github.com/br4tech/auth-nex/internal/model"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

type gormDBWithError struct {
	*gorm.DB
	err error
}

func (g *gormDBWithError) Error() error {
	return g.err
}

func TestCompanyRepository_CreateCompany(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock.NewMockIDatabase(ctrl)
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
		TenantID: 1,
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
		mockGormDB := &gorm.DB{}

		mockDB.EXPECT().Create(companyModel).Return(mockGormDB, nil)

		result, err := repo.CreateCompany(company)
		if err != nil {
			t.Fatalf("CreateCompany returned an error: %v", err)
		}

		if result == nil {
			t.Fatal("CreateCompany returned a nil result")
		}
	})

	t.Run("Error", func(t *testing.T) {
		expectedError := errors.New("database error")

		mockDB.EXPECT().Create(companyModel).Return(nil, expectedError)

		_, err := repo.CreateCompany(company)

		if err == nil {
			t.Fatal("CreateCompany did not return an error")
		}

		if err != expectedError {
			t.Fatalf("CreateCompany returned unexpected error: %v", err)
		}
	})
}
