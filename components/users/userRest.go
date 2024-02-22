package users

import (
	bn "github.com/kk/attendance_management/bean"
)

func AddUser(email string, role int, password string) error {
	var user bn.User
	user.Email = email
	user.Role = role
	user.Password = password
	err := AddUserSvc(&user)

	if err != nil {

		errMsg := make(map[string]string, 0)
		errMsg["error"] = err.Error()
	}
	return nil
}
