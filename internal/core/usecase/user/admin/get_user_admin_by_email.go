package admin

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type GetUserAdminByEmailUseCase struct {
	userRepository port.IUserRepository
}

func NewGetUserAdminByEmailUseCase(userRepository port.IUserRepository) port.IGetUserByEmailUseCase {
	return &GetUserAdminByEmailUseCase{
		userRepository: userRepository,
	}
}

func (uc *GetUserAdminByEmailUseCase) Execute(email string) (*domain.User, error) {
	user, err := uc.userRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
