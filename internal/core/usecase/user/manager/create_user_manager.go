package manager

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type CreateUserManagerUseCase struct {
	userRepository port.IUserRepository
}

func NewCreateUserManagerUseCase(userRepository port.IUserRepository) port.ICreateUserUsecase {
	return &CreateUserManagerUseCase{
		userRepository: userRepository,
	}
}

func (uc *CreateUserManagerUseCase) Execute(user *domain.User) error {
	user.Role = "manager"

	_, err := uc.userRepository.Create(user)
	if err != nil {
		return err
	}

	return nil
}
