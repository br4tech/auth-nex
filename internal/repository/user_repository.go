package repository

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db port.IDatabase
}

func NewUserRepository(db port.IDatabase) port.IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindUserByEmail(email string) (*domain.User, error) {
	var user model.User

	if err := r.db.Where("email=?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return user.ToDomain(), nil
}

func (r *UserRepository) CreateUser(user *domain.User) (*domain.User, error) {
	userModel := new(model.User)
	userModel.FromDomain(user)

	hashedPassword, err := hashPassword(userModel.Password)

	if err != nil {
		return nil, err
	}

	userModel.Password = hashedPassword

	_, err = r.db.Create(userModel)
	if err != nil {
		return nil, err
	}

	return userModel.ToDomain(), nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
