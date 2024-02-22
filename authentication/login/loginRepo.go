package login

import (
	// "fmt"

	bean "github.com/kk/attendance_management/bean"
	"github.com/kk/attendance_management/dataBase"
)

func LoginRepo(useremail string) (error, string) {

	db := dataBase.Connect()
	defer db.Close()

	var getpassword bean.User
	var mypass string

	err := db.Model(&getpassword).Column("password").Where("email=?", useremail).Select(&mypass)

	return err, mypass

}
