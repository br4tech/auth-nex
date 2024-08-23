package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserValidation(t *testing.T) {
	testCases := []struct {
		name      string
		user      *User
		validator func(User) error // Validador específico
		expectErr bool
	}{
		{
			name: "Valid System User",
			user: NewUser(
				1, "João", "joao@example.com", stringPtr("12345678901"), "senha123", "", 1, "system", intPtr(1),
			),
			validator: ValidateUserSystem, // Usando seu validador
			expectErr: false,
		},
		{
			name: "Invalid System User (Missing CPF)",
			user: NewUser(
				1, "Maria", "maria@example.com", nil, "senha123", "", 1, "system", intPtr(1),
			),
			validator: ValidateUserSystem, // Usando seu validador
			expectErr: true,
		},
		// ... (demais casos de teste para ClientUser, etc.)
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.validator(*tc.user)
			if tc.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// Funções auxiliares para criar ponteiros para string e int
func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}
