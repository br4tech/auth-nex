package repository

import "github.com/br4tech/auth-nex/internal/auth/model"

// AuthRepository define métodos para interagir com a autenticação.
type AuthRepository interface {
	FindUserByUsername(username string, tenantID uint) (*model.User, error)
	CreateUser(user *model.User) error
}
