package services

import (
	"database/sql"
	"fmt"
	"golang-default/models"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	db             *sql.DB
	sessionService *SessionService
}

func NewAuthService(db *sql.DB, session *SessionService) *AuthService {
	return &AuthService{
		db:             db,
		sessionService: session,
	}
}

func (s *AuthService) Login(email, password string) (map[string]interface{}, error) {
	user, err := s.getUserByEmail(email, password)
	if err != nil {
		return nil, err
	}

	// buat token pakai sessionService
	token, err := s.sessionService.CreateToken(user.ID)
	if err != nil {
		return nil, err
	}

	user.Password = "" // jangan return password

	return map[string]interface{}{
		"user":  user,
		"token": token,
	}, nil
}

// Fungsi helper untuk ambil user & cek password
func (s *AuthService) getUserByEmail(email, password string) (*models.LoginRequest, error) {
	var user models.LoginRequest
	err := s.db.QueryRow(
		`SELECT id, name, email, password FROM users WHERE email = ?`,
		email,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("invalid email or password")
	} else if err != nil {
		return nil, fmt.Errorf("failed to query user: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	return &user, nil
}
