package repository

import (
	"context"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
)

type PermissionRepository struct {
	db port.IDatabase[model.Profile]
}

func NewPermissionRepository(db port.IDatabase[model.Profile]) port.IPermissionRepository {
	return &PermissionRepository{db: db}
}

func (r *PermissionRepository) FindProfileByName(name string) (*domain.Profile, error) {
	profile, err := r.db.FindOne(context.Background(), "name=?", name)
	if err != nil {
		return nil, err
	}

	return profile.ToDomain(), nil
}

func (r *PermissionRepository) CreateProfile(role *domain.Profile) (*domain.Profile, error) {
	profileModel := new(model.Profile)
	profileModel.FromDomain(role)

	_, err := r.db.Create(context.Background(), profileModel)
	if err != nil {
		return nil, err
	}

	return profileModel.ToDomain(), nil
}
