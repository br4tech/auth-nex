package repository

import (
	"github.com/br4tech/auth-nex/internal/infra/model"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) IAuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) Authenticate(username string, tenantID uint) (*model.User, error) {
	var user model.User

	if err := r.db.Where("username=?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *AuthRepository) Register(userModel *model.User) error {
	user := model.User{
		UserName: userModel.UserName,
		Password: userModel.Password,
	}

	r.db.Create(&user)
	return nil
}
