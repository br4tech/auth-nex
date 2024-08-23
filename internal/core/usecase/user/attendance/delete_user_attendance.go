package attendance

import "github.com/br4tech/auth-nex/internal/core/port"

type DeleteUserAttendanceUsecase struct {
	userRepository port.IUserRepository
}

func NewDeleteUserAttendanceUsecase(userRepository port.IUserRepository) *DeleteUserAttendanceUsecase {
	return &DeleteUserAttendanceUsecase{
		userRepository: userRepository,
	}
}

func (uc *DeleteUserAttendanceUsecase) Execute(id int) error {
	return uc.userRepository.Delete(id)
}
