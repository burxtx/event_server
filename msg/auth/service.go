package auth

import (
	"context"

	"github.com/burxtx/car/users"
)

// Service is the interface that provides auth methods.
type Service interface {
	Register(ctx context.Context, username, password string) (*users.User, error)
	// Login(username, password string) bool
	// Logout() error
	// LoadUser(ID users.UserID) (User, error)
	// ConfirmEmail(email string) error
	// ConfirmPhone(number string) error
	// ConfirmPasswordReset(passwd string) error
}

type service struct {
	repo      users.UserRepository
	passwdSec users.PasswordSecurity
}

func New(repo users.UserRepository, pwdSec users.PasswordSecurity) Service {
	return &service{
		repo:      repo,
		passwdSec: pwdSec,
	}
}

func (s *service) Register(ctx context.Context, username, password string) (*users.User, error) {
	secPwd, err := s.passwdSec.Encrypt(password)
	if err != nil {
		return nil, err
	}
	a := users.NewAccount(username, secPwd)
	u := users.NewUser(a)
	err = s.repo.Create(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// User is a real model for auth view
type User struct {
	UserID   string
	NickName string
}
