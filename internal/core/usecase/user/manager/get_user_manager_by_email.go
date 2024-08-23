package manager

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type GetUserManagerByEmailUseCase struct {
	userRepository port.IUserRepository
}

func NewGetUserManagerByEmailUseCase(userRepository port.IUserRepository) *GetUserManagerByEmailUseCase {
	return &GetUserManagerByEmailUseCase{
		userRepository: userRepository,
	}
}

func (uc *GetUserManagerByEmailUseCase) Execute(email string) (*domain.User, error) {
	filter := map[string]interface{}{"email": email}
	user, err := uc.userRepository.FindBy(filter)
	if err != nil {
		return nil, err
	}

	return user, nil
}
