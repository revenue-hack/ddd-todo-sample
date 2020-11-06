package tododm

import "github.com/revenue-hack/ddd-todo-sample/src/core/domain/userdm"

type TodoDomainService struct {
	todoRepo TodoRepository
}

func NewTodoDomainService(todoRepo TodoRepository) *TodoDomainService {
	return &TodoDomainService{todoRepo: todoRepo}
}

func (service *TodoDomainService) CanConfirmTodoByUserID(userID userdm.UserID) bool {
	c, err := service.todoRepo.CountConfirmStatusByUserID(userID)
	if err != nil {
		return false
	}
	return c != 5
}

func (service *TodoDomainService) CanTodoByUserID(userID userdm.UserID) bool {
	c, err := service.todoRepo.CountTodoStatusByUserID(userID)
	if err != nil {
		return false
	}
	return c != 3
}
