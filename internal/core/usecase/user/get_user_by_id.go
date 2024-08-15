package user

import "github.com/br4tech/auth-nex/internal/core/port"

type GetUserByIdUseCase struct {
	userRepo port.IUserRepository
}

func NewGetUserByIdUseCase(userRepo port.IUserRepository) *GetUserByIdUseCase {
	return &GetUserByIdUseCase{
		userRepo: userRepo,
	}
}

func (uc *GetUserByIdUseCase) Execute(id int) error {
	return nil
}
