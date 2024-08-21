package admin

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type UpdateUserAdminUsecase struct {
	userRepository port.IUserRepository
}

func NewUpdateUserAdminUsecase(userRepository port.IUserRepository) *UpdateUserAdminUsecase {
	return &UpdateUserAdminUsecase{
		userRepository: userRepository,
	}
}

func (uc *UpdateUserAdminUsecase) Execute(user *domain.User) error {
	err := uc.userRepository.Update(user)
	if err != nil {
		return err
	}

	return nil
}
