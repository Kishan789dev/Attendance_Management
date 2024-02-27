package tokenvalid

import (
	"github.com/go-pg/pg"
	"github.com/kk/attendance_management/bean"
	"github.com/kk/attendance_management/dataBase"
)

type AuthenticationRepo interface {
	ValidateTokenAndGetEmailRepo(email string) error
}
type AuthenticationRepoImpl struct {
	database dataBase.DataBase
}

func NewAuthenticationRepo(database dataBase.DataBase) *AuthenticationRepoImpl {
	return &AuthenticationRepoImpl{
		database: database,
	}
}

func (impl *AuthenticationRepoImpl) ValidateTokenAndGetEmailRepo(email string) error {
	db := impl.database.Connect()

	err := db.Model(&bean.User{}).Where("email=?", email).Select()
	if err == pg.ErrNoRows {

		return err
	}
	return nil
}
