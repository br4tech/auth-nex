package model

import (
	"testing"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestPartner_ToDomain(t *testing.T) {
	partner := Partner{
		Participation: 33.33,
		UserId:        123,
		CompanyId:     456,
	}

	domainPartner := partner.ToDomain()

	assert.Equal(t, partner.Participation, domainPartner.Participation)
}

func TestPartner_FromDomain(t *testing.T) {
	domainPartner := domain.Partner{
		Participation: 33.33,
	}

	partner := Partner{}
	partner.FromDomain(&domainPartner)

	assert.Equal(t, domainPartner.Participation, partner.Participation)
}
