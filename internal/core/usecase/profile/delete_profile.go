package profile

import "github.com/br4tech/auth-nex/internal/core/port"

type DeleteProfileUseCase struct {
	profileRepo port.IProfileRepository
}

func NewDeleteProfileUseCase(profileRepo port.IProfileRepository) *DeleteProfileUseCase {
	return &DeleteProfileUseCase{
		profileRepo: profileRepo,
	}
}

func (uc *DeleteProfileUseCase) Execute(id int) error {
	return nil
}
