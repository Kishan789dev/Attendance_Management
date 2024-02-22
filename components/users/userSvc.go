package users

import bn "github.com/kk/attendance_management/bean"

func AddUserSvc(user *bn.User) error {
	return AddUserRepo(user)
}
