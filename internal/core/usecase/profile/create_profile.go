package profile

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
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
	profileModel := new(model.Profile)
	profileModel.FromDomain(profile)

	return uc.profileRepo.Create(profileModel)
}
