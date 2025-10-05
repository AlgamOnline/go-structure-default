package services

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type SessionService struct {
	SecretKey     string
	ExpireMinutes int
}

func NewSessionService(secret string) *SessionService {
	// Ambil expire time dari ENV
	expStr := os.Getenv("SESSION_EXPIRE_MINUTES")
	expMinutes, err := strconv.Atoi(expStr)
	if err != nil || expMinutes <= 0 {
		expMinutes = 1 // default 1 menit
	}

	return &SessionService{
		SecretKey:     secret,
		ExpireMinutes: expMinutes,
	}
}

// Buat token
func (s *SessionService) CreateToken(userID int64) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Minute * time.Duration(s.ExpireMinutes)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.SecretKey))
}

// Validasi token
func (s *SessionService) ValidateToken(tokenStr string) (bool, int, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.SecretKey), nil
	})

	if err != nil {
		return false, 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := int(claims["user_id"].(float64))
		return true, userID, nil
	}

	return false, 0, errors.New("invalid token")
}
