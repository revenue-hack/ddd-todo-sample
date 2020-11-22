package todoapp

type ListTodoItem struct {
	ID           string
	Title        string
	Contents     string
	ReferenceURL string
	UserName     string
}

type TodoQueryService interface {
	FindAll() ([]*ListTodoItem, error)
}
