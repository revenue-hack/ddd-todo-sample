package vo

import (
	"reflect"
	"regexp"

	"golang.org/x/xerrors"
)

type Email string

var (
	emailFormat = `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`
	emailRegExp = regexp.MustCompile(emailFormat)
)

const emailMaxLength = 256

func NewEmail(email string) (Email, error) {
	if len(email) == 0 {
		return Email(""), xerrors.New("email must not be empty")
	}

	if len(email) > emailMaxLength {
		return Email(""), xerrors.Errorf("email must less than %d: %s", emailMaxLength, email)
	}

	if ok := emailRegExp.MatchString(email); !ok {
		return Email(""), xerrors.Errorf("invalid email format. email is %s", email)
	}

	return Email(email), nil
}

func (e Email) Value() string {
	return string(e)
}

func (e Email) Equals(e2 Email) bool {
	return reflect.DeepEqual(e, e2)
}
