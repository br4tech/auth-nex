package admin

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type CreateUserAdminUsecase struct {
	userReposiotry port.IUserRepository
}

func NewCreateUsereAdminUseCase(userReposiotry port.IUserRepository) *CreateUserAdminUsecase {
	return &CreateUserAdminUsecase{
		userReposiotry: userReposiotry,
	}
}

func (uc *CreateUserAdminUsecase) Execute(user *domain.User) error {
	user.Role = "admin"

	_, err := uc.userReposiotry.Create(user)
	if err != nil {
		return err
	}

	return nil
}
