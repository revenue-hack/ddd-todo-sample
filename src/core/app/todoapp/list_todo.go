package todoapp

type ListTodoApp struct {
	todoQueryService TodoQueryService
}

func NewListTodoApp(todoQS TodoQueryService) *ListTodoApp {
	return &ListTodoApp{
		todoQueryService: todoQS,
	}
}

func (app *ListTodoApp) Exec() ([]*ListTodoItem, error) {
	return app.todoQueryService.FindAll()
}
