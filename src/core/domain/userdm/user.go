package userdm

import (
	"github.com/revenue-hack/ddd-todo-sample/src/core/domain/vo"
	"golang.org/x/exp/utf8string"
	"golang.org/x/xerrors"
)

type User struct {
	id       UserID
	userName string
	email    vo.Email
	password vo.Password
}

const (
	maxUserNameLength = 20
)

func NewUser(id UserID, userName string, email vo.Email, password vo.Password) (*User, error) {
	if userName == "" {
		return nil, xerrors.New("user name must be not empty")
	}
	if !utf8string.NewString(userName).IsASCII() {
		return nil, xerrors.Errorf("user name must be ASCII: %s", userName)
	}
	if len(userName) > maxUserNameLength {
		return nil, xerrors.Errorf("user name must be less and equal than %d", maxUserNameLength)
	}
	return &User{
		id:       id,
		userName: userName,
		email:    email,
		password: password,
	}, nil
}

func (u *User) ID() UserID {
	return u.id
}
func (u *User) UserName() string {
	return u.userName
}
func (u *User) Email() vo.Email {
	return u.email
}
func (u *User) Password() vo.Password {
	return u.password
}
func (u *User) Equals(u2 *User) bool {
	return u.id.Equals(u2.id)
}
