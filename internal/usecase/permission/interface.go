package usecase

type (
	IPermissionUseCase interface {
		CreateRole(name string) error
	}
)
