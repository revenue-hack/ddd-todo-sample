package tododm

import "golang.org/x/xerrors"

type Status uint

const (
	ICEBOX Status = iota
	TODO
	CONFIRM
	DONE
)

func NewStatusByPrimitive(s uint) (Status, error) {
	if uint(ICEBOX) > s {
		return 0, xerrors.Errorf("status is cast error: %d", s)
	}
	if uint(DONE) < s {
		return 0, xerrors.Errorf("status is cast error: %d", s)
	}

	return Status(s), nil
}

func (s Status) Value() uint {
	return uint(s)
}

func (s Status) IsConfirm() bool {
	return s == CONFIRM
}
func (s Status) IsToDO() bool {
	return s == TODO
}
