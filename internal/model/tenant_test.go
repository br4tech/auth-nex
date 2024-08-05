package model

import (
	"testing"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestTenant_ToDomain(t *testing.T) {
	tenant := Tenant{
		Name: "Tenant Exemplo",
		Companies: []Company{
			{LegalName: "Empresa A", TradeName: "Nome A", TenantID: 123},
			{LegalName: "Empresa B", TradeName: "Nome B", TenantID: 123},
		},
		Users: []User{
			{Name: "João", Email: "joao@example.com"},
			{Name: "Maria", Email: "maria@example.com"},
		},
	}

	domainTenant := tenant.ToDomain()

	assert.Equal(t, tenant.Name, domainTenant.Name)
	assert.Equal(t, len(tenant.Companies), len(domainTenant.Companies))
	assert.Equal(t, len(tenant.Users), len(domainTenant.Users))
}

func TestTenant_FromDomain(t *testing.T) {
	domainTenant := domain.Tenant{
		Name: "Tenant Exemplo",
		Companies: []domain.Company{
			{LegalName: "Empresa A", TradeName: "Nome A"},
			{LegalName: "Empresa B", TradeName: "Nome B"},
		},
		Users: []domain.User{
			{Name: "João", Email: "joao@example.com"},
			{Name: "Maria", Email: "maria@example.com"},
		},
	}

	tenant := Tenant{}
	tenant.FromDomain(&domainTenant)

	assert.Equal(t, domainTenant.Name, tenant.Name)
	assert.Equal(t, len(domainTenant.Companies), len(tenant.Companies))
	assert.Equal(t, len(domainTenant.Users), len(tenant.Users))
}
