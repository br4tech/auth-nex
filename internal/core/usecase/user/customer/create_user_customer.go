package customer

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type CreateUserCustomerUsecase struct {
	userRepository port.IUserRepository
}

func NewCreateUserCustomerUsecase(userRepository port.IUserRepository) port.ICreateUserUsecase {
	return &CreateUserCustomerUsecase{
		userRepository: userRepository,
	}
}

func (uc *CreateUserCustomerUsecase) Execute(user *domain.User) error {
	user.Role = "customer"

	_, err := uc.userRepository.Create(user)
	if err != nil {
		return err
	}

	return nil
}
