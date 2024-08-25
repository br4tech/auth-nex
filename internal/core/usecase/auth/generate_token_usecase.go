package auth

import (
	"time"

	"github.com/br4tech/auth-nex/internal/model"
	"github.com/golang-jwt/jwt/v4"
)

type GenerateTokenUseCase struct {
	jwtKey []byte
}

func NewGenerateTokenUseCase(jwtKey []byte) *GenerateTokenUseCase {
	return &GenerateTokenUseCase{
		jwtKey: jwtKey,
	}
}

func (uc *GenerateTokenUseCase) Execute(id int) (string, error) {
	claims := &model.Claims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(), // Token expira em 15 minutos
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(uc.jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
