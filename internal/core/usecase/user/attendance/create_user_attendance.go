package attendance

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type CreateUserAttendanceUsecase struct {
	userRepository port.IUserRepository
}

func NewCreateUserAttendanceUsecase(userRepository port.IUserRepository) *CreateUserAttendanceUsecase {
	return &CreateUserAttendanceUsecase{
		userRepository: userRepository,
	}
}

func (uc *CreateUserAttendanceUsecase) Execute(user *domain.User) error {
	user.Role = "attendance"

	_, err := uc.userRepository.Create(user)
	if err != nil {
		return nil
	}

	return nil
}
