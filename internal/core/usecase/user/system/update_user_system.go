package system

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type UpdateUserSystemUseCase struct {
	userRepository port.IUserRepository
}

func NewUpdateUserSystemUseCase(userRepository port.IUserRepository) *UpdateUserSystemUseCase {
	return &UpdateUserSystemUseCase{
		userRepository: userRepository,
	}
}

func (uc *UpdateUserSystemUseCase) Execute(user *domain.User) error {
	return uc.userRepository.Update(user)
}
