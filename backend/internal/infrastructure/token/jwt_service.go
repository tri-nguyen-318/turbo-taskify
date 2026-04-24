package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"userId"`
	jwt.RegisteredClaims
}

type Service struct {
	accessSecret  []byte
	refreshSecret []byte
	accessTTL     time.Duration
	refreshTTL    time.Duration
}

func NewService(accessSecret, refreshSecret string, accessTTL, refreshTTL time.Duration) *Service {
	return &Service{
		accessSecret:  []byte(accessSecret),
		refreshSecret: []byte(refreshSecret),
		accessTTL:     accessTTL,
		refreshTTL:    refreshTTL,
	}
}

func (s *Service) GenerateAccessToken(userID string) (string, error) {
	return s.sign(userID, s.accessSecret, s.accessTTL)
}

func (s *Service) GenerateRefreshToken(userID string) (string, error) {
	return s.sign(userID, s.refreshSecret, s.refreshTTL)
}

func (s *Service) ValidateAccessToken(raw string) (*Claims, error) {
	return s.parse(raw, s.accessSecret)
}

func (s *Service) ValidateRefreshToken(raw string) (*Claims, error) {
	return s.parse(raw, s.refreshSecret)
}

func (s *Service) sign(userID string, secret []byte, ttl time.Duration) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
}

func (s *Service) parse(raw string, secret []byte) (*Claims, error) {
	tok, err := jwt.ParseWithClaims(raw, &Claims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := tok.Claims.(*Claims)
	if !ok || !tok.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
