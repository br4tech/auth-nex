package attendance

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type GetUserAttendanceByEmailUsecase struct {
	userRepository port.IUserRepository
}

func NewGetUserAttendanceByEmailUsecase(userRepository port.IUserRepository) port.IGetUserByEmailUseCase {
	return &GetUserAttendanceByEmailUsecase{
		userRepository: userRepository,
	}
}

func (uc *GetUserAttendanceByEmailUsecase) Execute(email string) (*domain.User, error) {
	user, err := uc.userRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
