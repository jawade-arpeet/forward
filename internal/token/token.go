package token

import (
	"fmt"
	"forward/internal/config"

	"time"

	"github.com/golang-jwt/jwt/v5"
)

const accessTokenExpiryDuration = 1 * time.Hour       // 1 hour
const refreshTokenExpiryDuration = 7 * 24 * time.Hour // 7 days

type AccessTokenPayload struct {
	Email string `json:"email"`
	Type  string `json:"type"`
	jwt.RegisteredClaims
}

type RefreshTokenPayload struct {
	Type string `json:"type"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(
	id string,
	accID string,
	email string,
	issuedAt time.Time,
) (string, error) {
	cfg := config.GetJWTConfig()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, AccessTokenPayload{
		ID:        id,
		Email:     email,
		Type:      "access",
		Issuer:    cfg.Issuer,
		Subject:   accID,
		Audience:  cfg.Audience,
		IssuedAt:  jwt.NewNumericDate(issuedAt),
		ExpiresAt: jwt.NewNumericDate(issuedAt.Add(accessTokenExpiryDuration)),
	})

	tokenString, err := token.SignedString([]byte(cfg.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateAccessToken(tokenString string) (*AccessTokenPayload, error) {
	cfg := config.GetJWTConfig()

	claims := &AccessTokenPayload{}

	parsed, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(parsed *jwt.Token) (any, error) {
			if parsed.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("unexpected signing method: %v", parsed.Header["alg"])
			}
			return []byte(cfg.Secret), nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		jwt.WithIssuer(cfg.Issuer),
		jwt.WithAudience(cfg.Audience...),
	)
	if err != nil {
		return nil, err
	}

	if parsed == nil || !parsed.Valid || claims.Type != "access" {
		return nil, fmt.Errorf("invalid access token")
	}

	return claims, nil
}

func GenerateRefreshToken(
	id string,
	issuedAt time.Time,
) (string, error) {
	cfg := config.GetJWTConfig()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, RefreshTokenPayload{
		ID:        id,
		Type:      "refresh",
		Issuer:    cfg.Issuer,
		Audience:  cfg.Audience,
		IssuedAt:  jwt.NewNumericDate(issuedAt),
		ExpiresAt: jwt.NewNumericDate(issuedAt.Add(refreshTokenExpiryDuration)),
	})

	tokenString, err := token.SignedString([]byte(cfg.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
