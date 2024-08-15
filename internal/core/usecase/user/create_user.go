package user

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
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
	userModel := new(model.User)
	userModel.FromDomain(user)

	return uc.userReposiotry.Create(userModel)
}
