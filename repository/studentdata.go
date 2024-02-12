package repository

import (
	"github.com/kk/attendance_management/bean"
	"github.com/kk/attendance_management/dataBase"
	"log"
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

func AddStudentService(student bean.Student) (*bean.Student, error) {
	db := dataBase.Connect()
	defer db.Close()

	if _, err := db.Model(&student).Insert(); err != nil {
		log.Println(err)
		return nil, err
	}

	return &student, nil

}
