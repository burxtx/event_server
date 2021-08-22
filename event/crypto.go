package event

import (
	"golang.org/x/crypto/bcrypt"
)

// PasswordSecurity provides password security
type PasswordSecurity interface {
	Encrypt(password string) (Password, error)
	Compare(secPassword, plainPassword string) error
}

type passwordSecurity struct {
}

func (s *passwordSecurity) Encrypt(passwd string) (Password, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.MinCost)
	if err != nil {
		return Password{passwdHash: string(hash)}, err
	}
	return Password{passwdHash: string(hash)}, nil
}

func (s *passwordSecurity) Compare(secPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(secPassword), []byte(plainPassword))
}

func NewPasswordSecurity() PasswordSecurity {
	return &passwordSecurity{}
}
