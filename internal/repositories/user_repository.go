package repositories

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
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

func (repo *userRepositoryImpl) Create(user *domain.User) error {
	userModel := new(model.User)
	userModel.FromDomain(user)

	return repo.db.Create(user).Error
}

func (repo *userRepositoryImpl) FindById(id int) (*domain.User, error) {
	var user model.User
	result := repo.db.First(&user, id)

	return user.ToDomain(), result.Error
}

func (repo *userRepositoryImpl) Update(user *domain.User) error {
	userModel := new(model.User)
	userModel.FromDomain(user)

	return repo.db.Save(userModel).Error
}

func (repo *userRepositoryImpl) Delete(id int) error {
	return repo.db.Delete(&model.User{}, id).Error
}
