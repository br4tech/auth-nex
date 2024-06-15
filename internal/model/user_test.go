package model

import (
	"testing"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestUser_ToDomain(t *testing.T) {
	user := User{
		Name:      "João Silva",
		Email:     "joao@example.com",
		CPF:       "123.456.789-00",
		Password:  "senha123",
		Phone:     "19 99999-9999",
		TenantId:  1,
		Role:      "system",
		ProfileId: 2,
	}

	domainUser := user.ToDomain()

	assert.Equal(t, user.Name, domainUser.Name)
	assert.Equal(t, user.Email, domainUser.Email)
	assert.Equal(t, user.CPF, *domainUser.CPF)
	assert.Equal(t, user.Password, domainUser.Password)
	assert.Equal(t, user.Phone, domainUser.Phone)
	assert.Equal(t, user.TenantId, domainUser.TenantId)
	assert.Equal(t, user.Role, domainUser.Role)
	assert.Equal(t, user.ProfileId, *domainUser.ProfileId)
}

func TestUser_FromDomain(t *testing.T) {
	cpfPtr := "123.456.789-00"
	profileIdPtr := 1

	domainUser := domain.User{
		Name:      "João Silva",
		Email:     "joao@example.com",
		CPF:       &cpfPtr,
		Password:  "senha123",
		Phone:     "19 99999-9999",
		TenantId:  1,
		Role:      "system",
		ProfileId: &profileIdPtr,
	}

	user := User{}
	user.FromDomain(&domainUser)

	assert.Equal(t, domainUser.Name, user.Name)
	assert.Equal(t, domainUser.Email, user.Email)
	assert.Equal(t, *domainUser.CPF, user.CPF)
	assert.Equal(t, domainUser.Password, user.Password)
	assert.Equal(t, domainUser.Phone, user.Phone)
	assert.Equal(t, domainUser.TenantId, user.TenantId)
	assert.Equal(t, domainUser.Role, user.Role)
	assert.Equal(t, *domainUser.ProfileId, user.ProfileId)
}
