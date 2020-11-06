package todoapp

import (
	"github.com/revenue-hack/ddd-todo-sample/src/core/domain/tododm"
	"github.com/revenue-hack/ddd-todo-sample/src/core/domain/userdm"
	"github.com/revenue-hack/ddd-todo-sample/src/core/domain/vo"
	"golang.org/x/xerrors"
)

type UpdateTodoApp struct {
	todoRepository tododm.TodoRepository
	userRepository userdm.UserRepository
}

func NewUpdateTodoApp(todoRepo tododm.TodoRepository, userRepo userdm.UserRepository) *UpdateTodoApp {
	return &UpdateTodoApp{
		todoRepository: todoRepo,
		userRepository: userRepo,
	}
}

type UpdateTodoRequest struct {
	ID            string
	Title         string
	Contents      string
	Status        uint
	AsigneeUserID string
	ReferenceURL  string
}

// ドメイン知識を外に操作されないようにする
func (app *UpdateTodoApp) Exec(req *UpdateTodoRequest) error {
	todoID, err := tododm.NewTodoIDWithPrivimite(req.ID)
	if err != nil {
		return err
	}

	todo, err := app.todoRepository.FindByID(todoID)
	if err != nil {
		return err
	}

	status, err := tododm.NewStatusByPrimitive(req.Status)
	if err != nil {
		return err
	}
	asigneeID, err := userdm.NewUserIDWithStr(req.AsigneeUserID)
	if err != nil {
		return err
	}

	todoDomainService := tododm.NewTodoDomainService(app.todoRepository)
	if status.IsConfirm() && !todoDomainService.CanConfirmTodoByUserID(asigneeID) {
		return xerrors.Errorf("user can't belongs to this todo: %v", asigneeID)
	}
	if status.IsToDO() && !todoDomainService.CanTodoByUserID(asigneeID) {
		return xerrors.Errorf("user can't belongs to this todo: %v", asigneeID)
	}

	url, err := vo.NewURL(req.ReferenceURL)
	if err != nil {
		return err
	}

	if err := todo.ChangeContents(req.Contents); err != nil {
		return err
	}
	if err := todo.ChangeTitle(req.Title); err != nil {
		return err
	}
	todo.ChangeReferenceURL(url)
	todo.ChangeStatus(status)

	if _, err := app.todoRepository.Update(todo); err != nil {
		return err
	}

	return nil
}
