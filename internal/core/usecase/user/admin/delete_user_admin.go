package admin

import "github.com/br4tech/auth-nex/internal/core/port"

type DeleteUserAdminUsecase struct {
	userRepository port.IUserRepository
}

func NewDeleteUserAdminUsecase(userRepository port.IUserRepository) port.IDeleteUserUsecase {
	return &DeleteUserAdminUsecase{
		userRepository: userRepository,
	}
}

func (uc *DeleteUserAdminUsecase) Execute(id int) error {
	err := uc.userRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
