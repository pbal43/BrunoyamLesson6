package server

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func newRegisterClaims(issuer, audience, subject, jti string, ttl time.Duration) jwt.RegisteredClaims {
	now := time.Now()
	return jwt.RegisteredClaims{
		Issuer:    issuer,
		Subject:   subject,
		Audience:  jwt.ClaimStrings{audience},
		ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
		IssuedAt:  jwt.NewNumericDate(now),
		ID:        jti,
	}
}

func NewAccessToken(userID string) (string, error) {
	rc := newRegisterClaims(
		"My-api",
		"my",
		userID,
		"access token",
		time.Hour,
	)
	claims := Claims{
		UserID:           userID,
		RegisteredClaims: rc,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret string"))
}
