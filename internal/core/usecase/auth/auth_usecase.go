package auth

import (
	"errors"
	"time"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/dto"
	"github.com/br4tech/auth-nex/internal/model"
	"github.com/br4tech/auth-nex/pkg/validator"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("chave_secreta_do_jwt")

type AuthUseCase struct {
	userRepository port.IUserRepository
}

func NewAuthUseCase(userRepository port.IUserRepository) port.IUserUseCase {
	return &AuthUseCase{userRepository: userRepository}
}

func (uc *AuthUseCase) Authenticate(userReq *dto.UserTokenDTO) (*string, error) {
	user, err := uc.userRepository.FindUserByEmail(userReq.Email)

	if err != nil {
		return nil, err
	}

	if !comparePassword(userReq.Password, user.Password) {
		return nil, errors.New("Incorrect password")
	}

	token, err := generateAccessToken(user.Email)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (uc *AuthUseCase) CreateUser(user *domain.User) (*domain.User, error) {
	userModel := &model.User{}

	errUser := copier.Copy(userModel, user)
	if errUser != nil {
		return nil, errUser
	}

	if err := validator.ValidateStruct(userModel); err != nil {
		return nil, err
	}
	uc.userRepository.CreateUser(userModel.ToDomain())

	return userModel.ToDomain(), nil
}

func (uc *AuthUseCase) ValidateAccessToken(tokenString string) (*model.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("Token invalid")
}

func generateAccessToken(email string) (string, error) {
	claims := &model.Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(), // Token expira em 15 minutos
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func comparePassword(plainPassword string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
