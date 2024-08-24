package system

import "github.com/br4tech/auth-nex/internal/core/port"

type DeleteUserSystemUseCase struct {
	userRepository port.IUserRepository
}

func NewDeleteUserSystemUseCase(userRepository port.IUserRepository) port.IDeleteUsecase {
	return &DeleteUserSystemUseCase{
		userRepository: userRepository,
	}
}

func (uc *DeleteUserSystemUseCase) Execute(id int) error {
	return uc.userRepository.Delete(id)
}
