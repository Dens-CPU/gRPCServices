package jwtdto

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type RefreshToken struct {
	ID         int
	Token      string
	Expire_at  time.Time
	IsRevoked  bool
	UserId     uuid.UUID
	Created_at time.Time
}

type PairToken struct {
	AccsesToken  string
	RefreshToken string
	Expire_at    time.Duration
}

type AccsessClaim struct {
	User_id string
	email   string
	jwt.RegisteredClaims
}

func NewRefreshToken(user_id uuid.UUID) *RefreshToken {
	return &RefreshToken{
		Token:     uuid.NewString(),
		Expire_at: time.Now().Add(2 * time.Hour),
		IsRevoked: false,
		UserId:    user_id,
	}
}

func NewPairToken(accsesToken string, refreshToken string, expire_at time.Duration) PairToken {
	return PairToken{
		AccsesToken:  accsesToken,
		RefreshToken: refreshToken,
		Expire_at:    expire_at,
	}
}

func NewAccsessClaim(user_id, email string) *AccsessClaim {
	return &AccsessClaim{
		User_id: user_id,
		email:   email,
	}
}
