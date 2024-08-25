package auth

import (
	"errors"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/core/validator"
	"github.com/br4tech/auth-nex/internal/dto"
)

var jwtKey = []byte("chave_secreta_do_jwt")

type AuthenticateUserUseCase struct {
	userRepository       port.IUserRepository
	generateTokenUseCase port.IGenerateTokenUseCase
}

func NewAuthenticateUserUseCase(userRepository port.IUserRepository) *AuthenticateUserUseCase {
	return &AuthenticateUserUseCase{
		userRepository: userRepository,
	}
}

func (uc *AuthenticateUserUseCase) Execute(userReq *dto.UserTokenDTO) (*string, error) {
	var user *domain.User
	var err error

	if userReq.Email != "" && userReq.Password != "" {
		user, err = uc.authenticateWithEmailAndPassword(userReq.Email)
		if err != nil {
			return nil, err
		}
		if !validator.ComparePassword(userReq.Password, user.Password) {
			return nil, errors.New("Incorrect password")
		}
	} else if userReq.Phone != "" {
		user, err = uc.authenticateWithCelular(userReq.Phone)
		if err != nil {
			return nil, err
		}
	}

	token, err := uc.generateTokenUseCase.Execute(user.Id)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (uc *AuthenticateUserUseCase) authenticateWithEmailAndPassword(email string) (*domain.User, error) {
	user, err := uc.userRepository.FindByEmail(email)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *AuthenticateUserUseCase) authenticateWithCelular(celular string) (*domain.User, error) {
	user, err := uc.userRepository.FindByEmail(celular)

	if err != nil {
		return nil, err
	}
	return user, nil
}
