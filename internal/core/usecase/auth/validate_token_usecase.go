package auth

import (
	"errors"

	"github.com/br4tech/auth-nex/internal/model"
	"github.com/golang-jwt/jwt"
)

type ValidateTokenUseCase struct {
	jwtKey []byte
}

func NewValidateTokenUseCase(jwtKey []byte) *ValidateTokenUseCase {
	return &ValidateTokenUseCase{
		jwtKey: jwtKey,
	}
}

func (uc *ValidateTokenUseCase) Execute(tokenString string) (*model.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("Token inválido")
}
