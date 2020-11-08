package repoimpl

import "github.com/revenue-hack/ddd-todo-sample/src/core/domain/userdm"

type UserRepoImpl struct{}

func NewUserRepoImpl() *UserRepoImpl {
	return &UserRepoImpl{}
}

func (repo *UserRepoImpl) Create(user *userdm.User) (*userdm.User, error) {
	return nil, nil
}
func (repo *UserRepoImpl) FindByID(userID userdm.UserID) (*userdm.User, error) {
	return nil, nil
}
