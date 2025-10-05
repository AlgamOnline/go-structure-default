package services

import (
	"database/sql"
	"fmt"
	"golang-default/models"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

// Create user
func (s *UserService) CreateUser(user models.UserData) (int64, error) {
	// ✅ 1. Validasi format email
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return 0, fmt.Errorf("invalid email format")
	}

	// ✅ 2. Cek apakah email sudah ada
	var exists bool
	err = s.db.QueryRow(`SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)`, user.Email).Scan(&exists)
	if err != nil {
		return 0, fmt.Errorf("failed to check email: %w", err)
	}
	if exists {
		return 0, fmt.Errorf("email already registered")
	}

	// ✅ 3. Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("failed to hash password: %w", err)
	}

	// ✅ 4. Insert user
	result, err := s.db.Exec(
		`INSERT INTO users (name, email, password) VALUES (?, ?, ?)`,
		user.Name,
		user.Email,
		string(hashedPassword),
	)
	if err != nil {
		return 0, fmt.Errorf("failed to insert user: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %w", err)
	}

	return id, nil
}

// Get user by ID
func (s *UserService) GetUserByID(id int64) (models.UserData, error) {
	var user models.UserData
	err := s.db.QueryRow(`SELECT id, name, email FROM users WHERE id = ?`, id).
		Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user not found")
		}
		return user, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}

// Update user
func (s *UserService) UpdateUser(user models.UserData) error {
	_, err := s.db.Exec(`UPDATE users SET name = ?, email = ? WHERE id = ?`, user.Name, user.Email, user.ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

// Delete user
func (s *UserService) DeleteUser(id int64) error {
	_, err := s.db.Exec(`DELETE FROM users WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
