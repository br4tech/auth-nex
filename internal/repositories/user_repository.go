package repositories

import (
	"errors"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository[T port.IModel] struct {
	db port.IDatabase[T]
}

func NewUserRepository[T port.IModel](db port.IDatabase[T]) port.IUserRepository {
	return &UserRepository[T]{db: db}
}

func (r *UserRepository[T]) FindByEmail(email string) (*domain.User, error) {
	userEntity, err := r.db.FindBy("email", email)
	if err != nil {
		return nil, err
	}

	userModel, ok := any(userEntity[0]).(*model.User)
	if !ok {
		return nil, errors.New("failed to convert entity to user model")
	}

	userDomain := userModel.ToDomain()

	return userDomain, nil
}

func (r *UserRepository[T]) FindByPhone(phone string) (*domain.User, error) {
	var userEntity T

	_, err := r.db.FindBy("phone", phone)
	if err != nil {
		return nil, err
	}

	userModel, ok := any(&userEntity).(*model.User)
	if !ok {
		return nil, errors.New("failed to convert entity to user model")
	}

	userDomain := userModel.ToDomain()

	return userDomain, nil
}

func (r *UserRepository[T]) Create(user *domain.User) (*domain.User, error) {
	userModel := new(model.User)
	userModel.FromDomain(user)

	hashedPassword, err := hashPassword(userModel.Password)

	if err != nil {
		return nil, err
	}

	userModel.Password = hashedPassword

	userEntity, ok := any(userModel).(T)
	if !ok {
		return nil, errors.New("entity type invalid to user")
	}

	userId, err := r.db.Create(userEntity)
	if err != nil {
		return nil, err
	}

	user.Id = userId

	return user, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
