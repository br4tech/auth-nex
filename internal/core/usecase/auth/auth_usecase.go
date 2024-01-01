package usecase

import (
	"time"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("chave_secreta_do_jwt")

type AuthUseCase struct {
	userRepository port.IUserRepository
}

func NewAuthUseCase(userRepository port.IUserRepository) port.IUserUseCase {
	return &AuthUseCase{userRepository: userRepository}
}

func (uc *AuthUseCase) Authenticate(username, password string, tenantID int) (*domain.User, error) {
	return nil, nil
}

func (uc *AuthUseCase) CreateUser(user *model.User) error {
	return nil
}

func (uc *AuthUseCase) GenerateAccessToken(user *model.User) (string, error) {
	claims := &model.Claims{
		UserName: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(), // Token expira em 15 minutos
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (uc *AuthUseCase) ValidateAccessToken(tokenString string) (int, error) {
	return 0, nil
}
