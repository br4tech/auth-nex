package authentication

import "github.com/br4tech/auth-nex/internal/core/port"

type AuthenticateUseCase struct {
	userRepo port.IUserRepository
}

func NewAuthenticateUseCase(userRepo port.IUserRepository) *AuthenticateUseCase {
	return &AuthenticateUseCase{
		userRepo: userRepo,
	}
}

func (uc *AuthenticateUseCase) Execute() error {
	return nil
}
