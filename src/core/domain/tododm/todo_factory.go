package tododm

import (
	"github.com/revenue-hack/ddd-todo-sample/src/core/domain/userdm"
	"github.com/revenue-hack/ddd-todo-sample/src/core/domain/vo"
)

// ICEBOXが必ず決まっている
func GenerateTodoWhenUnCreated(title, contents string, asigneeUserID userdm.UserID, referenceURL vo.URL) (*Todo, error) {
	return newTodo(NewTodoID(), title, ICEBOX, contents, asigneeUserID, referenceURL)
}
