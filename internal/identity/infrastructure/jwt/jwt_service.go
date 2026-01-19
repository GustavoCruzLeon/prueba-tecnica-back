package jwt

import (
	"errors"
	"time"

	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type JWTService struct {
	secretKey string
}

func NewJWTService() *JWTService {
	godotenv.Load()
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "default-secret-key-for-development"
	}
	return &JWTService{secretKey: secret}
}

func (s *JWTService) GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

func (s *JWTService) ValidateToken(tokenStr string) (uint, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return 0, errors.New("invalid token")
	}

	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("user_id not found in token")
	}

	return uint(userID), nil
}
