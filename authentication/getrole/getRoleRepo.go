package getrole

import (
	"fmt"

	bn "github.com/kk/attendance_management/bean"
	"github.com/kk/attendance_management/dataBase"
)

func GetRoleRepo(email string) (error, int) {

	db := dataBase.Connect()
	defer db.Close()

	var role int

	err := db.Model(&bn.User{}).Column("role").Where("email=?", email).Select(&role)
	if err != nil {
		return fmt.Errorf("error while getting role reason:%s", err), role
	}
	return nil, role
}
