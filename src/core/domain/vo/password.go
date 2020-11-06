package vo

import (
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/xerrors"
)

type Password string

func NewPassword(pass string) (Password, error) {
	if pass == "" {
		return "", xerrors.New("password must be not empty")
	}

	if l := len(pass); l > 30 || l < 12 {
		return "", xerrors.New("password must be from 12 to 30 characters")
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	if err != nil {
		return "", xerrors.New("can't generate hash password")
	}
	return Password(string(hashPass)), nil
}

func (p Password) Value() string {
	return string(p)
}
