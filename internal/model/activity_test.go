package model

import (
	"testing"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestActivity_ToDomain(t *testing.T) {
	activity := Activity{
		CNAE:        "1234567",
		Description: "Descrição da atividade",
		CompanyID:   1,
	}

	domainActivity := activity.ToDomain()
	assert.Equal(t, activity.CNAE, domainActivity.CNAE)
	assert.Equal(t, activity.Description, domainActivity.Description)
	assert.Equal(t, activity.CompanyID, domainActivity.CompanyID)
}

func TestActivity_FromDomain(t *testing.T) {
	domainActivity := domain.Activity{
		CNAE:        "1234567",
		Description: "Descrição da atividade",
		CompanyID:   1,
	}

	activity := Activity{}
	activity.FromDomain(&domainActivity)
	assert.Equal(t, domainActivity.CNAE, activity.CNAE)
	assert.Equal(t, domainActivity.Description, activity.Description)
	assert.Equal(t, domainActivity.CompanyID, activity.CompanyID)
}
