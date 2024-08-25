package auth

import (
	"testing"
	"time"

	"github.com/br4tech/auth-nex/internal/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestGenerateTokenUseCase_Execute(t *testing.T) {
	jwtKey := []byte("chave_secreta_do_jwt")
	generateTokenUseCase := NewGenerateTokenUseCase(jwtKey)

	t.Run("Success", func(t *testing.T) {
		userId := 1

		tokenString, err := generateTokenUseCase.Execute(userId)

		assert.NoError(t, err)
		assert.NotEmpty(t, tokenString)

		token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		assert.NoError(t, err)
		assert.NotNil(t, token)

		claims, ok := token.Claims.(*model.Claims)
		assert.True(t, ok)
		assert.NotNil(t, claims)

		assert.Equal(t, userId, claims.Id)
		assert.True(t, claims.ExpiresAt > time.Now().Unix())
		assert.Equal(t, time.Now().Unix(), claims.IssuedAt)
	})
}
