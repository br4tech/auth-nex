package customer

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type GetUserCustomerByPhoneUseCase struct {
	userRepository port.IUserRepository
}

func NewGetUserCustomerByPhoneUseCase(userRepository port.IUserRepository) *GetUserCustomerByPhoneUseCase {
	return &GetUserCustomerByPhoneUseCase{
		userRepository: userRepository,
	}
}

func (uc *GetUserCustomerByPhoneUseCase) Execute(phone string) (*domain.User, error) {
	filter := map[string]interface{}{"phone": phone}

	user, err := uc.userRepository.FindBy(filter)
	if err != nil {
		return nil, err
	}

	return user, nil
}
