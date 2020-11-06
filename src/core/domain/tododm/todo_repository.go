package tododm

import "github.com/revenue-hack/ddd-todo-sample/src/core/domain/userdm"

type TodoRepository interface {
	Create(todo *Todo) (TodoID, error)
	Update(todo *Todo) (TodoID, error)
	FindByID(id TodoID) (*Todo, error)
	CountConfirmStatusByUserID(userID userdm.UserID) (uint, error)
	CountTodoStatusByUserID(userID userdm.UserID) (uint, error)
}
