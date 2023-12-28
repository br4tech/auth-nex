package usecase

import (
	"time"

	"github.com/br4tech/auth-nex/internal/infra/model"
	"github.com/br4tech/auth-nex/internal/infra/repository"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("chave_secreta_do_jwt")

type AuthUseCase struct {
	authRepository repository.IAuthRepository
}

func NewAuthUseCase(authRepository repository.IAuthRepository) IAuthUseCase {
	return &AuthUseCase{authRepository: authRepository}
}

func (uc *AuthUseCase) Authenticate(username, password string, tenantID int) (*model.User, error) {
	return nil, nil
}

func (uc *AuthUseCase) Register(user *model.User) error {
	return nil
}

func (uc *AuthUseCase) GenerateAccessToken(user *model.User) (string, error) {
	claims := &model.Claims{
		UserName: user.UserName,
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
