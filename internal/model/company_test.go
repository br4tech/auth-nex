package model

import (
	"testing"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestCompany_ToDomain(t *testing.T) {
	company := Company{
		LegalName:         "Nome Legal LTDA",
		TradeName:         "Nome Fantasia",
		Document:          "12345678901234",
		StateRegistration: "123456789",
		Address: Address{
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
		Partners: []Partner{
			{
				CompanyId:     1,
				Participation: 50,
				UserId:        1,
			},
		},
	}

	domainCompany := company.ToDomain()
	assert.Equal(t, company.LegalName, domainCompany.LegalName)
	assert.Equal(t, company.TradeName, domainCompany.TradeName)
	assert.Equal(t, company.Document, domainCompany.Document)
	assert.Equal(t, company.StateRegistration, domainCompany.StateRegistration)
	assert.Equal(t, company.Address.ToDomain(), &domainCompany.Address)
	assert.Equal(t, company.Type, domainCompany.Type)
	assert.Equal(t, uint(company.TenantId), domainCompany.TenantId)
	assert.Equal(t, company.Schema, domainCompany.Schema)
	assert.Equal(t, len(company.Partners), len(domainCompany.Partners))
}

func TestCompany_FromDomain(t *testing.T) {
	domainCompany := domain.Company{
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
	}

	company := Company{}
	company.FromDomain(&domainCompany)

	assert.Equal(t, domainCompany.LegalName, company.LegalName)
	assert.Equal(t, domainCompany.TradeName, company.TradeName)
	assert.Equal(t, domainCompany.Document, company.Document)
	assert.Equal(t, domainCompany.StateRegistration, company.StateRegistration)

	assert.NotNil(t, company.Address)
	assert.Equal(t, domainCompany.Address.Street, company.Address.Street)
	assert.Equal(t, domainCompany.Address.Number, company.Address.Number)
	assert.Equal(t, domainCompany.Address.Complement, company.Address.Complement)
	assert.Equal(t, domainCompany.Address.District, company.Address.District)
	assert.Equal(t, domainCompany.Address.City, company.Address.City)
	assert.Equal(t, domainCompany.Address.State, company.Address.State)
	assert.Equal(t, domainCompany.Address.ZipCode, company.Address.ZipCode)

	assert.Equal(t, domainCompany.Type, company.Type)
	assert.Equal(t, int(domainCompany.TenantId), company.TenantId)
	assert.Equal(t, domainCompany.Schema, company.Schema)
}
