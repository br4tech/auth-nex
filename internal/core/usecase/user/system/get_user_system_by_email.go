package system

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type GetUserSystemByEmailUseCase struct {
	userRepository port.IUserRepository
}

func NewGetUserSystemByEmailUseCase(userRepository port.IUserRepository) port.IGetUserByEmailUseCase {
	return &GetUserSystemByEmailUseCase{
		userRepository: userRepository,
	}
}

func (uc *GetUserSystemByEmailUseCase) Execute(email string) (*domain.User, error) {
	filter := map[string]interface{}{"email": email}
	user, err := uc.userRepository.FindBy(filter)
	if err != nil {
		return nil, err
	}

	return user, nil
}
