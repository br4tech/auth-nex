package repository

import "github.com/br4tech/auth-nex/internal/core/domain"

type MessageRepository struct{}

func NewMessageRepository() MessageRepository {
	return MessageRepository{}
}

func (r *MessageRepository) PushNotification() (*domain.Notification, error) {
	return nil, nil
}
