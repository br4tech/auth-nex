package repository

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) port.IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Authenticate(name string, tenantID uint) (*domain.User, error) {
	var user model.User

	if err := r.db.Where("name=?", name).First(&user).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *UserRepository) CreateUser(userModel *model.User) error {
	user := model.User{
		Name:     userModel.Name,
		Password: userModel.Password,
	}

	r.db.Create(&user)
	return nil
}
