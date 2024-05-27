package model

import (
	"testing"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestRole_ToDomain(t *testing.T) {
	role := Role{
		Name: "admin",
	}

	domainRole := role.ToDomain()

	assert.Equal(t, role.Name, domainRole.Name)
}

func TestRole_FromDomain(t *testing.T) {
	domainRole := domain.Role{
		Name: "admin",
	}

	role := Role{}
	role.FromDomain(&domainRole)

	assert.Equal(t, domainRole.Name, role.Name)
}
