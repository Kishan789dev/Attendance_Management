package getrole

import (
	"fmt"

	bn "github.com/kk/attendance_management/bean"
	"github.com/kk/attendance_management/dataBase"
)

type GetroleRepo interface {
	GetRoleRepo(email string) (error, int)
}
type GetroleRepoImpl struct {
	database dataBase.DataBase
}

func NewGetRole(database dataBase.DataBase) *GetroleRepoImpl {
	return &GetroleRepoImpl{

		database: database,
	}
}

func (impl *GetroleRepoImpl) GetRoleRepo(email string) (error, int) {

	db := impl.database.Connect()
	// defer db.Close()

	var role int

	err := db.Model(&bn.User{}).Column("role").Where("email=?", email).Select(&role)
	if err != nil {
		return fmt.Errorf("error while getting role reason:%s", err), role
	}
	return nil, role
}
