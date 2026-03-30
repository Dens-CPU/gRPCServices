package jwt

import "time"

type PairToken struct {
	AccsessToken string
	RefreshToken string
	Expire_at    time.Duration
}

func NewPairToken(accsesToken string, refreshToken string, expire_at time.Duration) PairToken {
	return PairToken{
		AccsessToken: accsesToken,
		RefreshToken: refreshToken,
		Expire_at:    expire_at,
	}
}
