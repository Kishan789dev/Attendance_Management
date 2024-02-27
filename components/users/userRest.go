package users

import (
	bn "github.com/kk/attendance_management/bean"
)

type UserRest interface {
	AddUser(email string, role int, password string) error
}
type UserRestImpl struct {
	usersvc UserSvc
}

func NewUserRestImpl(usersvc UserSvc) *UserRestImpl {
	return &UserRestImpl{
		usersvc: usersvc,
	}
}

func (impl *UserRestImpl) AddUser(email string, role int, password string) error {
	var user bn.User
	user.Email = email
	user.Role = role
	user.Password = password
	err := impl.usersvc.AddUserSvc(&user)

	if err != nil {

		errMsg := make(map[string]string, 0)
		errMsg["error"] = err.Error()
	}
	return nil
}
