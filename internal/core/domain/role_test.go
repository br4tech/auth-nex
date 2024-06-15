package domain

import (
	"testing"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
)

func TestNewRole(t *testing.T) {
	t.Run("validação bem-sucedida", func(t *testing.T) {
		role := NewRole("admin")
		validate := validator.New()
		err := validate.Struct(role)
		assert.NoError(t, err)
	})

	t.Run("nome ausente", func(t *testing.T) {
		role := NewRole("")
		validate := validator.New()
		err := validate.Struct(role)
		assert.Error(t, err)
		assert.True(t, IsRequiredFieldMissing(err))
	})

	t.Run("nome correto", func(t *testing.T) {
		expectedRole := &Role{
			Name: "admin",
		}
		role := NewRole("admin")
		assert.Equal(t, expectedRole, role)
	})
}
