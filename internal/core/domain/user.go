package domain

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Roles []Role
}
