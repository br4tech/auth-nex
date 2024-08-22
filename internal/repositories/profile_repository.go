package repositories

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
	"gorm.io/gorm"
)

type profileRepositoryImpl struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) port.IProfileRepository {
	return &profileRepositoryImpl{
		db: db,
	}
}

func (repo *profileRepositoryImpl) Create(profile *domain.Profile) (*domain.Profile, error) {
	profileModel := new(model.Profile)
	profileModel.FromDomain(profile)

	if err := repo.db.Create(profileModel).Error; err != nil {
		return nil, err
	}

	return profileModel.ToDomain(), nil
}

func (repo *profileRepositoryImpl) FindById(id int) (*domain.Profile, error) {
	var profile model.Profile
	result := repo.db.First(&profile, id)

	return profile.ToDomain(), result.Error
}

func (repo *profileRepositoryImpl) Upate(profile *domain.Profile) (*domain.Profile, error) {
	profileModel := new(model.Profile)
	profileModel.FromDomain(profile)

	if err := repo.db.Save(profileModel).Error; err != nil {
		return nil, err
	}

	return profileModel.ToDomain(), nil
}

func (repo *profileRepositoryImpl) Delete(id int) error {
	return repo.db.Delete(&model.Profile{}, id).Error
}
