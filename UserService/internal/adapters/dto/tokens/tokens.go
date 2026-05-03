package tokensdto

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type RefreshToken struct {
	ID        int
	Token     string
	ExpireAt  time.Time
	IsRevoked bool
	UserId    uuid.UUID
	CreatedAt time.Time
	UpdateAt  time.Time
}

type PairToken struct {
	AccessToken  string
	RefreshToken string
	ExpireAt     time.Time
}

type AccessClaim struct {
	UserId string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type InputTokens struct {
	AccsesToken  string
	RefreshToken string
}

func NewRefreshToken(userId uuid.UUID) *RefreshToken {
	return &RefreshToken{
		Token:     uuid.NewString(),
		ExpireAt:  time.Now().Add(5 * 24 * time.Hour),
		IsRevoked: false,
		UserId:    userId,
	}
}

func NewPairToken(accessToken string, refreshToken string, expireAt time.Time) PairToken {
	return PairToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpireAt:     expireAt,
	}
}

func NewAccessClaim(userId, email, role string) *AccessClaim {
	return &AccessClaim{
		UserId: userId,
		Email:  email,
		Role:   role,
	}
}

func NewInputTokens(accessToken string, refreshToken string) InputTokens {
	return InputTokens{
		AccsesToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
