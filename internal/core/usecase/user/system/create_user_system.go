package system

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type CreateUserSystemUseCase struct {
	userRepository port.IUserRepository
}

func NewCreateUserSytemUseCase(userRepository port.IUserRepository) port.ICreateUserUsecase {
	return &CreateUserSystemUseCase{
		userRepository: userRepository,
	}
}

func (uc *CreateUserSystemUseCase) Execute(user *domain.User) error {
	user.Role = "system_admin"

	_, err := uc.userRepository.Create(user)
	if err != nil {
		return err
	}
	return nil
}
