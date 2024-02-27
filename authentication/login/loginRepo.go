package login

import (
	bean "github.com/kk/attendance_management/bean"
	"github.com/kk/attendance_management/dataBase"
)

type LoginRepo interface {
	LoginRepo(useremail string) (error, string)
}
type LoginRepoImpl struct {
	database dataBase.DataBase
}

func NewLoginRepo(database dataBase.DataBase) *LoginRepoImpl {
	return &LoginRepoImpl{
		database: database,
	}
}

func (impl *LoginRepoImpl) LoginRepo(useremail string) (error, string) {

	// fmt.Println(dataBase)
	db := impl.database.Connect()

	var getpassword bean.User
	var mypass string

	err := db.Model(&getpassword).Column("password").Where("email=?", useremail).Select(&mypass)

	return err, mypass

}
