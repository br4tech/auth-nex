package usecase

type TenantUseCase interface {
	CreateTenant(name string) error
}
