package userapp

import (
	"github.com/revenue-hack/ddd-todo-sample/src/core/domain/userdm"
	"github.com/revenue-hack/ddd-todo-sample/src/core/domain/vo"
)

type CreateUserApp struct {
	userRepository userdm.UserRepository
}

func NewCreateUserApp(userRepo userdm.UserRepository) *CreateUserApp {
	return &CreateUserApp{
		userRepository: userRepo,
	}
}

type CreateUserRequest struct {
	UserName string
	Email    string
	Password string
}

type CreateUserResponse struct {
	ID string
}

func (app *CreateUserApp) Exec(req *CreateUserRequest) (*CreateUserResponse, error) {
	email, err := vo.NewEmail(req.Email)
	if err != nil {
		return nil, err
	}
	password, err := vo.NewPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user, err := userdm.NewUser(userdm.NewUserID(), req.UserName, email, password)
	if err != nil {
		return nil, err
	}

	createdUser, err := app.userRepository.Create(user)
	if err != nil {
		return nil, err
	}
	return &CreateUserResponse{ID: createdUser.ID().String()}, nil
}
