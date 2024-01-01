package dto

import "github.com/br4tech/auth-nex/internal/core/domain"

type Notification struct {
	Message string `json:"message"`
}

func (dto Notification) ToDomain() *domain.Notification {
	return &domain.Notification{
		Message: dto.Message,
	}
}

func (dto Notification) FromDomain(domain *domain.Notification) {
	dto.Message = domain.Message
}
