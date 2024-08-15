package user

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type CreateUserUseCase struct {
	userReposiotry port.IUserRepository
}

func NewCreateUserUseCase(userReposiotry port.IUserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		userReposiotry: userReposiotry,
	}
}

func (uc *CreateUserUseCase) Execute(user *domain.User) error {
	return uc.userReposiotry.Create(user)
}
