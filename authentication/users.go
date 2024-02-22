package authentication

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	bn "github.com/kk/attendance_management/bean"
	"github.com/kk/attendance_management/dataBase"
)

func GetRole(w http.ResponseWriter, r *http.Request) (int, error) {

	email, err := ValidateTokenAndGetEmail(w, r)
	if err != nil {
		json.NewEncoder(w).Encode("user is unauthorised")
		return 0, err

	}
	var usr bn.User
	db := dataBase.Connect()
	defer db.Close()

	var role int

	err = db.Model(&usr).Column("role").Where("email=?", email).Select(&role)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return 0, err
	}
	return role, err

}

func GetRoletemp(w http.ResponseWriter, r *http.Request, tokenstr string) (int, error) {

	email, err := ValidateTokenAndGetEmailtemp(w, r, tokenstr)
	if err != nil {
		json.NewEncoder(w).Encode("user is unauthorised")
		return 0, err

	}
	var usr bn.User
	db := dataBase.Connect()
	defer db.Close()
	var role int
	fmt.Println(db)
	err = db.Model(&usr).Column("role").Where("email=?", email).Select(&role)
	if err != nil {
		log.Println("Error in users.go", err)
		w.WriteHeader(http.StatusBadRequest)
		return 0, err
	}
	return role, err

}
