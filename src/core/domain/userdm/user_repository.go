package userdm

type UserRepository interface {
	Create(user *User) (User, error)
}
