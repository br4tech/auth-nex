package repositories

import (
	"errors"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
)

type ProfileRepository[T port.IModel] struct {
	db port.IDatabase[T]
}

func NewProfileRepository[T port.IModel](db port.IDatabase[T]) port.IProfileRepository {
	return &ProfileRepository[T]{db: db}
}

func (r *ProfileRepository[T]) FindByName(name string) (*domain.Profile, error) {
	profileEntity, err := r.db.FindBy("name", name)
	if err != nil {
		return nil, err
	}

	profileModel, ok := any(profileEntity[0]).(*model.Profile)
	if !ok {
		return nil, errors.New("failed to convert entity to profile model")
	}

	profileDomain := profileModel.ToDomain()

	return profileDomain, nil

}

func (r *ProfileRepository[T]) Create(profile *domain.Profile) (*domain.Profile, error) {
	profileModel := new(model.Profile)
	profileModel.FromDomain(profile)

	profileEntty, ok := any(profileModel).(T)
	if !ok {
		return nil, errors.New("entity type invalid to profile")
	}

	profileId, err := r.db.Create(profileEntty)
	if err != nil {
		return nil, err
	}

	profile.Id = profileId

	return profile, nil
}
