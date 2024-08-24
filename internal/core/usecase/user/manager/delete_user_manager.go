package manager

import "github.com/br4tech/auth-nex/internal/core/port"

type DeleteUserManagerUseCase struct {
	userRepository port.IUserRepository
}

func NewDeleteUserManagerUseCase(userRepository port.IUserRepository) port.IDeleteUsecase {
	return &DeleteUserManagerUseCase{
		userRepository: userRepository,
	}
}

func (uc *DeleteUserManagerUseCase) Execute(id int) error {
	return uc.userRepository.Delete(id)
}
