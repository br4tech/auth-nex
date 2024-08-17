package profile

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type CreateProfileUsecase struct {
	profileRepo port.IProfileRepository
}

func NewCreateProfileUseCase(profileRepo port.IProfileRepository) *CreateProfileUsecase {
	return &CreateProfileUsecase{
		profileRepo: profileRepo,
	}
}

func (uc *CreateProfileUsecase) Execute(profile *domain.Profile) error {
	return uc.profileRepo.Create(profile)
}
