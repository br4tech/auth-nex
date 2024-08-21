package profile

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type GetProfileByIdUseCase struct {
	profileRepo port.IProfileRepository
}

func NewGetProfileByIdUseCase(profileRepo port.IProfileRepository) *GetProfileByIdUseCase {
	return &GetProfileByIdUseCase{
		profileRepo: profileRepo,
	}
}

func (uc *GetProfileByIdUseCase) Execute(id int) (*domain.Profile, error) {
	profile, err := uc.profileRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	return profile, nil
}
