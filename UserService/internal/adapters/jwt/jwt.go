package jwt

import (
	"time"

	userconfig "github.com/DencCPU/gRPCServices/UserService/config"
	jwtdto "github.com/DencCPU/gRPCServices/UserService/internal/adapters/dto/tokens"
	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	secret          []byte
	accsessTokenTTL time.Duration
}

func NewJWT(cfg userconfig.JWT) *JWT {
	secret := cfg.Secret
	ttl := cfg.TTL
	return &JWT{secret: []byte(secret), accsessTokenTTL: time.Minute * time.Duration(ttl)}
}

func (j *JWT) CreateAccsesToken(user_id string, email string) (string, time.Duration, error) {
	accsessClaim := jwtdto.NewAccsessClaim(user_id, email)
	accsessClaim.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.accsessTokenTTL)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	accsessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accsessClaim).SignedString(j.secret)
	if err != nil {
		return "", 0, err
	}
	return accsessToken, j.accsessTokenTTL, nil
}
