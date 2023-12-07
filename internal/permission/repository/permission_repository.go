package repository

import "github.com/br4tech/auth-nex/internal/permission/model"

type PermissionRepository interface {
	FindRoleByName(name string) (*model.Role, error)
	CreateRole(role *model.Role) error
}
