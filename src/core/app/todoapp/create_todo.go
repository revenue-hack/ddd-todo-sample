package todoapp

import (
	"github.com/revenue-hack/ddd-todo-sample/src/core/domain/tododm"
	"github.com/revenue-hack/ddd-todo-sample/src/core/domain/userdm"
	"github.com/revenue-hack/ddd-todo-sample/src/core/domain/vo"
)

type CreateTodoApp struct {
	todoRepository tododm.TodoRepository
	userRepository userdm.UserRepository
}

func NewCreateTodoApp(todoRepo tododm.TodoRepository, userRepo userdm.UserRepository) *CreateTodoApp {
	return &CreateTodoApp{
		todoRepository: todoRepo,
		userRepository: userRepo,
	}
}

type CreateTodoRequest struct {
	Title         string
	Contents      string
	AsigneeUserID string
	ReferenceURL  string
}

type CreateTodoResponse struct {
	ID       string
	Title    string
	Status   uint
	Contents string
}

// ドメイン知識を外に操作されないようにする
func (app *CreateTodoApp) Exec(req *CreateTodoRequest) (*CreateTodoResponse, error) {
	asigneeID, err := userdm.NewUserIDWithStr(req.AsigneeUserID)
	if err != nil {
		return nil, err
	}
	url, err := vo.NewURL(req.ReferenceURL)
	if err != nil {
		return nil, err
	}

	todo, err := tododm.GenerateTodoWhenUnCreated(req.Title, req.Contents, asigneeID, url)
	if err != nil {
		return nil, err
	}

	if _, err := app.todoRepository.Create(todo); err != nil {
		return nil, err
	}

	return &CreateTodoResponse{
		ID:       todo.ID().String(),
		Title:    todo.Title(),
		Contents: todo.Contents(),
		Status:   todo.Status().Value(),
	}, nil
}
