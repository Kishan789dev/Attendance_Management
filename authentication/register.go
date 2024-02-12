package authentication

import (
	"encoding/json"
	"log"
	"net/http"

	bn "github.com/kk/attendance_management/bean"
	"github.com/kk/attendance_management/dataBase"
)

func AddUser(email string, role int, password string) error {
	var users bn.User
	users.Email = email
	users.Role = role
	users.Password = password

	db := dataBase.Connect()
	defer db.Close()

	_, err := db.Model(&users).Insert()
	if err != nil {
		log.Fatal(err, "add user fun")
		return err
	}
	return nil
}

func Register(w http.ResponseWriter, r *http.Request) {
	var userdetails bn.User

	_ = json.NewDecoder(r.Body).Decode(&userdetails)
	if userdetails.Role == 3 {
		err := AddUser(userdetails.Email, userdetails.Role, userdetails.Password)
		if err != nil {
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("only principle can register")
		return
	}
	// }

}
