package model

import (
	"testing"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestAddress_ToDomain(t *testing.T) {
	address := Address{
		Street:     "Rua Exemplo",
		Number:     "123",
		Complement: "Apto 456",
		District:   "Bairro Legal",
		City:       "Cidade Grande",
		State:      "SP",
		ZipCode:    "12345-678",
	}

	domainAddress := address.ToDomain()
	assert.Equal(t, address.Street, domainAddress.Street)
	assert.Equal(t, address.Number, domainAddress.Number)
	assert.Equal(t, address.Complement, domainAddress.Complement)
	assert.Equal(t, address.District, domainAddress.District)
	assert.Equal(t, address.City, domainAddress.City)
	assert.Equal(t, address.State, domainAddress.State)
	assert.Equal(t, address.ZipCode, domainAddress.ZipCode)
}

func TestAddress_FromDomain(t *testing.T) {
	domainAddress := domain.Address{
		Street:     "Rua Exemplo",
		Number:     "123",
		Complement: "Apto 456",
		District:   "Bairro Legal",
		City:       "Cidade Grande",
		State:      "SP",
		ZipCode:    "12345-678",
	}

	address := Address{}
	address.FromDomain(&domainAddress)
	assert.Equal(t, domainAddress.Street, address.Street)
	assert.Equal(t, domainAddress.Number, address.Number)
	assert.Equal(t, domainAddress.Complement, address.Complement)
	assert.Equal(t, domainAddress.District, address.District)
	assert.Equal(t, domainAddress.City, address.City)
	assert.Equal(t, domainAddress.State, address.State)
	assert.Equal(t, domainAddress.ZipCode, address.ZipCode)
}
