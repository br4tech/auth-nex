package repositories

import (
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) port.IUserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (repo *userRepositoryImpl) Create(user *model.User) error {
	return repo.db.Create(user).Error
}

func (repo *userRepositoryImpl) FindById(id int) (*model.User, error) {
	var user model.User
	result := repo.db.First(&user, id)

	return &user, result.Error
}

func (repo *userRepositoryImpl) Update(user *model.User) error {
	return repo.db.Save(user).Error
}

func (repo *userRepositoryImpl) Delete(id int) error {
	return repo.db.Delete(&model.User{}, id).Error
}
