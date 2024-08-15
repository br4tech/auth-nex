package user

import "github.com/br4tech/auth-nex/internal/core/port"

type DeleteUserUseCase struct {
	userRepo port.IUserRepository
}

func NewDeleteUserUseCase(userRepo port.IUserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		userRepo: userRepo,
	}
}

func (uc *DeleteUserUseCase) Execute(id int) error {
	return nil
}
