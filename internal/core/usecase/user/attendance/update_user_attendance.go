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

func (uc *UpdateUserAttendanceUsecase) Execute(user *domain.User) (*domain.User, error) {
	user, err := uc.userRepository.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
