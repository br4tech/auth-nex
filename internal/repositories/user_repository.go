package repositories

import (
	"errors"
	"fmt"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/core/validator"
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

func (repo *userRepositoryImpl) Create(user *domain.User) (*domain.User, error) {
	userModel := new(model.User)
	userModel.FromDomain(user)

	hashedPassword := validator.HashPassword(userModel.Password)
	if hashedPassword == "" {
		return nil, errors.New("Error hash password")
	}

	user.Password = hashedPassword

	if err := repo.db.Create(userModel).Error; err != nil {
		return nil, err
	}

	return userModel.ToDomain(), nil
}

func (repo *userRepositoryImpl) FindById(id int) (*domain.User, error) {
	var user model.User
	result := repo.db.First(&user, id)

	return user.ToDomain(), result.Error
}

func (repo *userRepositoryImpl) FindByEmail(email string) (*domain.User, error) {
	var user model.User
	result := repo.db.Where("email=?", email).First(&user)

	return user.ToDomain(), result.Error
}

func (repo *userRepositoryImpl) FindByPhone(phone string) (*domain.User, error) {
	var user model.User
	result := repo.db.Where("phone=?", phone).First(&user)

	return user.ToDomain(), result.Error
}

func (repo *userRepositoryImpl) FindBy(filter map[string]interface{}) (*domain.User, error) {
	var user model.User

	query := repo.db
	for field, value := range filter {
		query = query.Where(fmt.Sprintf("%s = ?", field), value)
	}

	if err := query.First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("erro ao buscar usuário: %w", err)
	}

	return user.ToDomain(), nil
}

func (repo *userRepositoryImpl) Update(user *domain.User) (*domain.User, error) {
	userModel := new(model.User)
	userModel.FromDomain(user)

	if err := repo.db.Save(userModel).Error; err != nil {
		return nil, err
	}

	return userModel.ToDomain(), nil
}

func (repo *userRepositoryImpl) Delete(id int) error {
	return repo.db.Delete(&model.User{}, id).Error
}
