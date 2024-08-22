package profile

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type UpdateProfileUseCase struct {
	profileRepo port.IProfileRepository
}

func NewUpdateProfile(profileRepo port.IProfileRepository) *UpdateProfileUseCase {
	return &UpdateProfileUseCase{
		profileRepo: profileRepo,
	}
}

func (uc *UpdateProfileUseCase) Execute(profile *domain.Profile) (*domain.Profile, error) {
	profile, err := uc.profileRepo.Upate(profile)
	if err != nil {
		return nil, err
	}

	return profile, nil
}
