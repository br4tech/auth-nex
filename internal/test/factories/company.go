package factories

import "github.com/br4tech/auth-nex/internal/core/domain"

func NewCompanyFactory() *domain.Company {
	return &domain.Company{
		Id:                1,
		LegalName:         "Empresa Legal Ltda.",
		TradeName:         "Empresa Legal",
		Document:          "12345678901234",
		StateRegistration: "123456789",
		Type:              "LTDA",
		TenantId:          1,
		Schema:            "app-0000",
		ParentCompanyId:   0, // Ou outro ID de empresa pai, se aplicável
		Address:           *NewAddressFactory(),
		Users:             []domain.User{*NewUserFactory("attendance")},
		Partners:          []domain.Partner{*NewPartnerFactory(10.0)},
		Activities:        []domain.Activity{*NewActivityFactory()},
	}
}
