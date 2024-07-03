package repository

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
)

type PermissionRepository struct {
	db port.IDatabase[model.Company]
}

func NewPermissionRepository(db port.IDatabase[model.Company]) port.IPermissionRepository {
	return &PermissionRepository{db: db}
}

func (r *PermissionRepository) FindProfileByName(name string) (*domain.Profile, error) {
	var profileModel model.Profile

	if err := r.db.Where("name=?", name).First(&profileModel).Error; err != nil {
		return nil, err
	}

	return profileModel.ToDomain(), nil
}

func (r *PermissionRepository) CreateProfile(role *domain.Profile) (*domain.Profile, error) {
	profileModel := new(model.Profile)
	profileModel.FromDomain(role)

	_, err := r.db.Create(profileModel)
	if err != nil {
		return nil, err
	}

	return profileModel.ToDomain(), nil
}
