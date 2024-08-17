package user

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type UpdateUserUseCase struct {
	userRepo port.IUserRepository
}

func NewUpdateUserUseCase(userRepo port.IUserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		userRepo: userRepo,
	}
}

func (uc *UpdateUserUseCase) Execute(user *domain.User) error {
	return uc.userRepo.Update(user)
}
