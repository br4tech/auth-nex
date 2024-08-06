package repositories

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/mock"
	"github.com/br4tech/auth-nex/internal/model"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestTenantRepository_CreateTenant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock.NewMockIDatabase[port.IModel](ctrl)
	repo := NewTenantRepository(mockDB)

	tenant := &domain.Tenant{
		Id:   123,
		Name: "Tenant Exemplo",
		Companies: []domain.Company{
			{LegalName: "Empresa A", TradeName: "Nome A", TenantID: 123},
			{LegalName: "Empresa B", TradeName: "Nome B", TenantID: 123},
		},
		Users: []domain.User{
			{Name: "João", Email: "joao@example.com"},
			{Name: "Maria", Email: "maria@example.com"},
		},
	}

	tenantModel := new(model.Tenant)
	tenantModel.FromDomain(tenant)

	t.Run("Success", func(t *testing.T) {
		mockGormDB := &gorm.DB{}

		mockDB.EXPECT().Create(tenantModel).Return(mockGormDB, nil)

		result, err := repo.Create(tenant)
		if err != nil {
			t.Fatalf("CreateTenant returned an error: %v", err)
		}

		if result == nil {
			t.Fatal("CreateTenant returned a nil result")
		}

		if result.Id != tenant.Id {
			t.Fatalf("Expected tenant ID %d, got %d", tenant.Id, result.Id)
		}
	})

	t.Run("Error", func(t *testing.T) {
		expectedError := errors.New("database error")
		mockDB.EXPECT().Create(tenantModel).Return(nil, expectedError)

		_, err := repo.Create(tenant)

		if err == nil {
			t.Fatal("CreateTenant did not return an error")
		}

		if err != expectedError {
			t.Fatalf("CreateTenant returned unexpected error: %v", err)
		}
	})
}
