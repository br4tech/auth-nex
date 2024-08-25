package customer

import "github.com/br4tech/auth-nex/internal/core/port"

type DeleteUserCustomerUsecase struct {
	userRepository port.IUserRepository
}

func NewDeleteUserCustomerUsecase(userRepository port.IUserRepository) port.IDeleteUserUsecase {
	return &DeleteUserCustomerUsecase{
		userRepository: userRepository,
	}
}

func (uc *DeleteUserCustomerUsecase) Execute(id int) error {
	return uc.userRepository.Delete(id)
}
