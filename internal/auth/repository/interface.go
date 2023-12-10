package repository

import "github.com/br4tech/auth-nex/internal/auth/model"

type (
	IAuthRepository interface {
		Authenticate(username, password string, tenantID uint) (*model.User, error)
		Register(user *model.User) error
	}
)
