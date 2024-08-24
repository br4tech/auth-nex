package system

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type UpdateUserSystemUseCase struct {
	userRepository port.IUserRepository
}

func NewUpdateUserSystemUseCase(userRepository port.IUserRepository) port.IUpdateUserUsecase {
	return &UpdateUserSystemUseCase{
		userRepository: userRepository,
	}
}

func (uc *UpdateUserSystemUseCase) Execute(user *domain.User) (*domain.User, error) {
	user, err := uc.userRepository.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
