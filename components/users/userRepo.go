package users

import (
	"fmt"

	bn "github.com/kk/attendance_management/bean"
	"github.com/kk/attendance_management/dataBase"
)

type UserRepo interface {
	AddUserRepo(user *bn.User) error
}
type UserRepoImpl struct {
	database dataBase.DataBase
}

func NewUserRepo(database dataBase.DataBase) *UserRepoImpl {
	return &UserRepoImpl{
		database: database,
	}
}

func (impl *UserRepoImpl) AddUserRepo(user *bn.User) error {
	db := impl.database.Connect()

	_, err := db.Model(user).Insert()
	return fmt.Errorf("error occured during inserting into db in user table reason:%s", err)

}
