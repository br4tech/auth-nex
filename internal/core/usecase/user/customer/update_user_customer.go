package customer

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type UpdateUserCustomerUsecase struct {
	userRepository port.IUserRepository
}

func NewUpdateUserCustomerUsecase(userRepository port.IUserRepository) *UpdateUserCustomerUsecase {
	return &UpdateUserCustomerUsecase{
		userRepository: userRepository,
	}
}

func (uc *UpdateUserCustomerUsecase) Execute(user *domain.User) error {
	return uc.userRepository.Update(user)
}
