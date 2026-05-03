package postgresdto

import (
	"time"

	userhash "github.com/DencCPU/gRPCServices/UserService/internal/adapters/hash"
	"github.com/google/uuid"
)

type UserDTO struct {
	ID           uuid.UUID
	Name         string
	Email        string
	HashPassword string
	Role         string
	CreatedAt    time.Time
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
	UserId     int
	CreatedAt  time.Time
}

type AuthUser struct {
	ID   string
	Role string
}

func NewUserDTO(name, email, password, role string) (*UserDTO, error) {
	hashPassword, err := userhash.HashPassword(password)
	if err != nil {
		return nil, err
	}
	dto := UserDTO{
		Name:         name,
		Email:        email,
		HashPassword: hashPassword,
		Role:         role,
	}
	return &dto, nil
}

func NewUpdatePassord(email, password string) *UpdatePassword {
	return &UpdatePassword{Email: email, Password: password}
}
