package domain

import (
	"testing"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
)

func TestNewActivity(t *testing.T) {
	t.Run("validação bem-sucedida", func(t *testing.T) {
		activity := NewActivity("1234567", "Descrição da atividade", 1)
		validate := validator.New()
		err := validate.Struct(activity)
		assert.NoError(t, err)
	})

	t.Run("campos obrigatórios ausentes", func(t *testing.T) {
		testCases := []struct {
			cnae        string
			description string
			companyID   int
		}{
			{"", "Descrição da atividade", 1},
			{"1234567", "", 1},
			{"1234567", "Descrição da atividade", 0},
		}

		validate := validator.New()
		for _, tc := range testCases {
			activity := NewActivity(tc.cnae, tc.description, tc.companyID)
			err := validate.Struct(activity)
			assert.Error(t, err)
			assert.True(t, IsRequiredFieldMissing(err))
		}
	})
}
