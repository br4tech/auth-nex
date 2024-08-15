package profile

import "github.com/br4tech/auth-nex/internal/core/port"

type GetProfileByIdUseCase struct {
	profileRepo port.IProfileRepository
}

func NewGetProfileByIdUseCase(profileRepo port.IProfileRepository) *GetProfileByIdUseCase {
	return &GetProfileByIdUseCase{
		profileRepo: profileRepo,
	}
}

func (uc *GetProfileByIdUseCase) Execute(id int) error {
	return nil
}
