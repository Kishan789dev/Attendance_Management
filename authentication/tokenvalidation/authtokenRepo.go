package tokenvalid

import (
	"github.com/go-pg/pg"
	"github.com/kk/attendance_management/bean"
	"github.com/kk/attendance_management/dataBase"
)

func ValidateTokenAndGetEmailRepo(email string) error {
	db := dataBase.Connect()
	defer db.Close()
	err := db.Model(&bean.User{}).Where("email=?", email).Select()
	if err == pg.ErrNoRows {

		return err
	}
	return nil
}
