package users

import (
	"fmt"

	bn "github.com/kk/attendance_management/bean"
	"github.com/kk/attendance_management/dataBase"
)

func AddUserRepo(user *bn.User) error {
	db := dataBase.Connect()
	defer db.Close()
	_, err := db.Model(user).Insert()
	return fmt.Errorf("error occured during inserting into db in user table reason:%s", err)

}
