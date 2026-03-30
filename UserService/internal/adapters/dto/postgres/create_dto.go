package postgresdto

import (
	"time"

	"github.com/google/uuid"
)

type UserDTO struct {
	ID         uuid.UUID
	Name       string
	Email      string
	Password   string
	Role       string
	Created_at time.Time
}

type UpdatePassword struct {
	Email     string
	Password  string
	Update_at time.Time
}

type RefreshToken struct {
	ID         int
	Token      string
	Expires_at time.Time
	IsRevoked  bool
	User_id    int
	Created_at time.Time
}

func NewUserDTO(name, email, password, role string) *UserDTO {
	dto := UserDTO{
		Name:     name,
		Email:    email,
		Password: password,
		Role:     role,
	}
	return &dto
}

func NewUpdatePassord(email, password string) *UpdatePassword {
	return &UpdatePassword{Email: email, Password: password}
}
