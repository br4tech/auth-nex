package model

import (
	"testing"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestUser_ToDomain(t *testing.T) {
	user := User{
		Name:          "João Silva",
		Email:         "joao@example.com",
		Password:      "senha123",
		CPF:           "123.456.789-00",
		Nationality:   "Brasileira",
		MaritalStatus: "Solteiro",
		Address: Address{
			Street:     "Rua Exemplo",
			Number:     "123",
			Complement: "Apto 456",
			District:   "Bairro Legal",
			City:       "Cidade Grande",
			State:      "SP",
			ZipCode:    "12345-678",
		},
		TenantId:  1,
		ProfileId: 2,
	}

	domainUser := user.ToDomain()

	assert.Equal(t, user.Name, domainUser.Name)
	assert.Equal(t, user.Email, domainUser.Email)
	assert.Equal(t, user.Password, domainUser.Password)
	assert.Equal(t, user.TenantId, domainUser.TenantId)
	assert.Equal(t, user.ProfileId, domainUser.ProfileId)
}

func TestUser_FromDomain(t *testing.T) {
	domainUser := domain.User{
		Name:      "João Silva",
		Email:     "joao@example.com",
		Password:  "senha123",
		TenantId:  1,
		ProfileId: 2,
	}

	user := User{}
	user.FromDomain(&domainUser)

	assert.Equal(t, domainUser.Name, user.Name)
	assert.Equal(t, domainUser.Email, user.Email)
	assert.Equal(t, domainUser.Password, user.Password)
	assert.Equal(t, domainUser.TenantId, user.TenantId)
	assert.Equal(t, domainUser.ProfileId, user.ProfileId)
}
