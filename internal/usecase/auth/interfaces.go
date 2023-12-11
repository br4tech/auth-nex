package usecase

import "github.com/br4tech/auth-nex/internal/infra/model"

type (
	IAuthUseCase interface {
		Authenticate(username, password string, tenantID int) (*model.User, error)
		Register(user *model.User) error
		GenerateAccessToken(user *model.User) (string, error)
		ValidateAccessToken(tokenString string) (int, error)
	}
)
