package usecase

import "github.com/br4tech/auth-nex/internal/auth/model"

type (
	IAuthUseCase interface {
		Authenticate(username, password string, tenantID uint) (*model.User, error)
		Register(user *model.User) error
		GenerateAccessToken(user *model.User) (string, error)
		ValidateAccessToken(tokenString string) (uint, error)
	}
)
