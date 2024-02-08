package repository

import (
	"log"

	"github.com/kk/attendance_management/bean"
	"github.com/kk/attendance_management/dataBase"
)

func GetStudentsRepo() ([]bean.Student, error) {

	db := dataBase.Connect()
	defer db.Close()

	var students []bean.Student

	err := db.Model(&students).Select()
	if err != nil {
		log.Println(err)
	}
	return students, err

}

func GetStudentRepo(id int) (bean.Student, error) {

	db := dataBase.Connect()
	defer db.Close()

	students := bean.Student{Sid: id}

	if err := db.Model(&students).Where("sid=?", id).Select(); err != nil {
		log.Println(err)
		return students, err
	}

	return students, nil
}

func AddStudentRepo() {
	var userdetails bean.Userdetails
		db := dataBase.Connect()
		defer db.Close()

		_ = json.NewDecoder(r.Body).Decode(&userdetails)

		student := bean.Student{Name: userdetails.Name, Address: userdetails.Address, Class: userdetails.Class, Email: userdetails.Email}
		// _ = json.NewDecoder(r.Body).Decode(&student)

		// student.Id = uuid.New().String()

		if _, err := db.Model(&student).Insert(); err != nil {
			log.Println(err)
			// json.NewEncoder(w).Encode("error is line no 77")

			w.WriteHeader(http.StatusBadRequest)
			return

}

func UpdateStudentRepo() {

}

func DeleteStudentRepo() {

}
