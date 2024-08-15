package repositories

import (
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

func (repo *profileRepositoryImpl) Create(profile *model.Profile) error {
	return repo.db.Create(profile).Error
}

func (repo *profileRepositoryImpl) FindById(id int) (*model.Profile, error) {
	var profile model.Profile
	result := repo.db.First(&profile, id)

	return &profile, result.Error
}

func (repo *profileRepositoryImpl) Upate(profile *model.Profile) error {
	return repo.db.Save(profile).Error
}

func (repo *profileRepositoryImpl) Delete(id int) error {
	return repo.db.Delete(&model.Profile{}, id).Error
}
