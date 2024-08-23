package factories

import "github.com/br4tech/auth-nex/internal/core/domain"

func NewPartnerFactory(participation float64) *domain.Partner {
	return &domain.Partner{
		Participation: participation,
		UserId:        1,
		User:          *NewUserFactory("admin"),
	}
}
