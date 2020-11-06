package tododm

import (
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type TodoID string

func NewTodoID() TodoID {
	return TodoID(uuid.New().String())
}

func (id TodoID) String() string {
	return string(id)
}

func NewTodoIDWithPrivimite(id string) (TodoID, error) {
        if id == "" {
                return "", xerrors.New("todo id must be not empty")
        }
        return TodoID(id), nil
}
