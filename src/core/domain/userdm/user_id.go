package userdm

import (
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type UserID string

func NewUserID() UserID {
	return UserID(uuid.New().String())
}

func NewUserIDWithStr(id string) (UserID, error) {
	if id == "" {
		return "", xerrors.New("user id  must be not empty")
	}
	return UserID(id), nil
}

func (id UserID) Equals(id2 UserID) bool {
	return string(id) == string(id2)
}

func (id UserID) String() string {
	return string(id)
}
