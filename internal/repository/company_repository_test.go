package repository

import (
	"context"
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
		mockDB.EXPECT().Create(gomock.Any(), gomock.Any()).
			DoAndReturn(func(ctx context.Context, value *model.Company) (*model.Company, error) {
				value.Id = 123
				return value, nil
			})

		repo := NewCompanyRepository(mockDB)

		createdCompany, err := repo.CreateCompany(company)

		if err != nil {
			t.Errorf("Erro inesperado: %v", err)
		}

		if createdCompany.Id != 123 {
			t.Errorf("ID do usuário criado incorreto. Esperado: %d, Obtido: %d", 123, createdCompany.Id)
		}
	})

	t.Run("Error", func(t *testing.T) {
		expectedError := errors.New("database error")

		mockDB.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, expectedError)

		_, err := repo.CreateCompany(company)

		if err == nil {
			t.Fatal("CreateCompany did not return an error")
		}

		if err != expectedError {
			t.Fatalf("CreateCompany returned unexpected error: %v", err)
		}
	})
}
