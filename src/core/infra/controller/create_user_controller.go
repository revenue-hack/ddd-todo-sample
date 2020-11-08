package controller

import (
	"github.com/revenue-hack/ddd-todo-sample/src/core/app/userapp"
	"github.com/revenue-hack/ddd-todo-sample/src/core/infra/repoimpl"
)

func Exec() error {
	userRepo := repoimpl.NewUserRepoImpl()
	app := userapp.NewCreateUserApp(userRepo)
	req := &userapp.CreateUserRequest{
		UserName: "username",
		Email:    "aaa@bbb.com",
		Password: "23jijrdrao",
	}
	if _, err := app.Exec(req); err != nil {
		return err
	}
	return nil
}
