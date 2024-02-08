package restHandler

import (
	"encoding/json"
	"log"
	"net/http"

	bn "github.com/kk/attendance_management/bean"
	"github.com/kk/attendance_management/dataBase"
)

func AddUser(w http.ResponseWriter, email string, role int, password string) {
	var users bn.User
	users.Email = email
	users.Role = role
	users.Password = password

	db := dataBase.Connect()
	defer db.Close()

	_, err := db.Model(&users).Insert()
	if err != nil {
		log.Fatal(err, "add user fun")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	var userdetails bn.User

	_ = json.NewDecoder(r.Body).Decode(&userdetails)
	// if userdetails.Role == 1 {
	// 	AddStudent(w, userdetails.Name, userdetails.Address, userdetails.Class, userdetails.Email)
	// 	AddUser(w, userdetails.Email, userdetails.Role, userdetails.Password)

	// } else if userdetails.Role == 2 {

	// 	AddTeacher(w, userdetails.Name, userdetails.Address, userdetails.Email)
	// 	AddUser(w, userdetails.Email, userdetails.Role, userdetails.Password)

	// } else {
	if userdetails.Role == 3 {

		AddUser(w, userdetails.Email, userdetails.Role, userdetails.Password)

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("only principle can register")
		return
	}
	// }

}
