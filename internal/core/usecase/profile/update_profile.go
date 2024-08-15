package profile

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
)

type UpdateProfileUseCase struct {
	profileRepo port.IProfileRepository
}

func NewUpdateProfile(profileRepo port.IProfileRepository) *UpdateProfileUseCase {
	return &UpdateProfileUseCase{
		profileRepo: profileRepo,
	}
}

func (uc *UpdateProfileUseCase) Execute(profile *domain.Profile) error {
	profileModel := new(model.Profile)
	profileModel.FromDomain(profile)

	return uc.profileRepo.Upate(profileModel)
}
