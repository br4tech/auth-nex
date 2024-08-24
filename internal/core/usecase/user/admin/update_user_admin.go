package admin

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type UpdateUserAdminUsecase struct {
	userRepository port.IUserRepository
}

func NewUpdateUserAdminUsecase(userRepository port.IUserRepository) port.IUpdateUserUsecase {
	return &UpdateUserAdminUsecase{
		userRepository: userRepository,
	}
}

func (uc *UpdateUserAdminUsecase) Execute(user *domain.User) (*domain.User, error) {
	user, err := uc.userRepository.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
