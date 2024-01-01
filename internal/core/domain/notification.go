package domain

type Notification struct {
	Message string `validate:"required"`
}

func NewNotification(message string) *Notification {
	return &Notification{
		Message: message,
	}
}
