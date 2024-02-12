package services

import (
	"github.com/kk/attendance_management/bean"
	repository "github.com/kk/attendance_management/repository"
)

func GetStudentsSvc() ([]bean.Student, error) {

	students, err := repository.GetStudentsRepo()
	return students, err
}

func GetStudentSvc(id int) (bean.Student, error) {

	students, err := repository.GetStudentRepo(id)
	if err != nil {
		return students, err
	}

	return students, nil
}

func AddStudentService(student bean.Student) (*bean.Student, error) {

	students, err := repository.AddStudentService(student)
	if err != nil {

		return students, err
	}

	return students, nil
}
