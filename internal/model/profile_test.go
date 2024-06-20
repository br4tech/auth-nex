package model

import (
	"testing"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestProfile_ToDomain(t *testing.T) {
	profile := Profile{
		Name: "admin",
	}

	domainProfile := profile.ToDomain()

	assert.Equal(t, profile.Name, domainProfile.Name)
}

func TestProfile_FromDomain(t *testing.T) {
	domainProfile := domain.Profile{
		Name: "admin",
	}

	profile := Profile{}
	profile.FromDomain(&domainProfile)

	assert.Equal(t, domainProfile.Name, profile.Name)
}
