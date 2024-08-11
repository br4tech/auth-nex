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

func TestTenantRepository_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock.NewMockIDatabase[port.IModel](ctrl)
	repo := NewTenantRepository(mockDB)

	tenant := &domain.Tenant{
		Id:   123,
		Name: "Tenant Exemplo",
		Companies: []domain.Company{
			{LegalName: "Empresa A", TradeName: "Nome A", TenantId: 123},
			{LegalName: "Empresa B", TradeName: "Nome B", TenantId: 123},
		},
		Users: []domain.User{
			{Name: "João", Email: "joao@example.com"},
			{Name: "Maria", Email: "maria@example.com"},
		},
	}

	tenantModel := new(model.Tenant)
	tenantModel.FromDomain(tenant)

	t.Run("Success", func(t *testing.T) {
		mockDB.EXPECT().Create(tenantModel).Return(tenant.Id, nil)

		result, err := repo.Create(tenant)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, tenant.Id, result.Id)
	})

	t.Run("Error", func(t *testing.T) {
		expectedError := errors.New("database error")
		mockDB.EXPECT().Create(tenantModel).Return(0, expectedError)

		_, err := repo.Create(tenant)

		assert.Error(t, err)
		assert.EqualError(t, err, "database error")
	})
}

func TestTenantRepository_FindByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock.NewMockIDatabase[model.Tenant](ctrl)
	repo := NewTenantRepository(mockDB)

	t.Run("Success", func(t *testing.T) {
		name := "APP-000000000"

		expectedTenants := model.Tenant{
			Name: name,
		}

		mockDB.EXPECT().FindBy("name=?", name).Return(&expectedTenants, nil)

		tenant, err := repo.FindByName(name)
		assert.NoError(t, err)
		assert.NotNil(t, tenant)
		assert.Equal(t, expectedTenants.Name, tenant.Name)
	})

	t.Run("TenantNotFound", func(t *testing.T) {
		name := "APP-000000001"
		mockDB.EXPECT().FindBy("name=?", name).DoAndReturn(func(f, v string) ([]*model.User, error) {
			return nil, errors.New("TenantNotFound")
		})

		tenant, err := repo.FindByName(name)
		assert.Error(t, err)
		assert.Nil(t, tenant)
	})
}
