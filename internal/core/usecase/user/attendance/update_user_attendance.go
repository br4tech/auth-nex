package attendance

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type UpdateUserAttendanceUsecase struct {
	userRepository port.IUserRepository
}

func NewUpdateUserAttendanceUsecase(userRepository port.IUserRepository) *UpdateUserAttendanceUsecase {
	return &UpdateUserAttendanceUsecase{
		userRepository: userRepository,
	}
}

func (uc *UpdateUserAttendanceUsecase) Execute(user *domain.User) error {
	return uc.userRepository.Update(user)
}
