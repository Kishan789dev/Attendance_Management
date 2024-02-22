package students

import (
	"fmt"
	"log"

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

func AddStudentService(student *bean.Student) error {

	err := repository.AddStudentService(student)
	if err != nil {

		return err
	}

	return nil
}

// punchin

func StudentAttendanceWithPunchData(sid int) (error, int) {
	aid, err := repository.StudentEntryPunchinEntryRepo(sid)
	if err != nil {
		return err, 0
	}

	err = repository.StudentEntryPunchinEntryTableRepo(aid)
	if err != nil {
		return err, 1
	}

	return nil, aid
}

func StudentPunchEntryInTable(aid int) (error, string) {
	pi_count, po_count, err := repository.StudentEntryOnlyPunchinSvc(aid)
	if err != nil {
		return err, ""
	}

	if pi_count <= po_count {

		err = repository.StudentPunchingSvc(aid)
		if err != nil {
			log.Println("lololololololololooloolol", err)

			return err, ""
		}
		return err, "punch in successful"

	} else {
		return err, "already punch in"

	}

}

func StudentEntryPunchinSvc(email string) (error, int, int) {
	err, typ, sid := repository.StudentEntryPunchinRepo(email)
	fmt.Println("jdsjsfjsjdsjsdjdsjdsjdsjdssdjjs", sid)
	if err != nil {
		log.Println("sidddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd")
		return err, typ, sid
	}

	return nil, typ, sid

}

// *********punchout
var aid int

func StudentEntryPunchOutSvc(email string) (error, int, int, int) {
	err, typ, sid, aid := repository.StudentEntryPunchOutRepo(email)
	if err != nil {

		return err, typ, sid, aid
	}

	return nil, 3, sid, aid

}

// func StudentAttendanceWithPunchOutData() (error, int) {
// 	// aid, err := repository.StudentEntryPunchinEntryRepo()
// 	// if err != nil {
// 	// 	return err, 0
// 	// }

// 	err := repository.StudentEntryPunchinEntryTableRepo(aid)
// 	if err != nil {
// 		return err, 0
// 	}

// 	return nil, aid
// }

func StudentPunchOutEntryInTable(aid int) (error, string) {
	pi_count, po_count, err := repository.StudentEntryOnlyPunchOutSvc(aid)
	if err != nil {
		return err, ""
	}

	if pi_count > po_count {

		err = repository.StudentPunchingOutRepo(aid)
		if err != nil {

			return err, ""
		}
		return nil, "Puchout successfully"

	} else {

		return nil, "you have already punch out"

	}

}

// Getstudents attendance

func GetStudentsAttendanceEmailSvc(email string) (error, int) {
	err, sid := repository.GetSid(email)
	if err != nil {
		return err, 0

	}
	return nil, sid

}
func FetchAttendanceFromDetailsSvc(studentattendance *bean.StudentAttendance) (*[]bean.StudentAttendancetemp, error) {
	return repository.FetchAttendanceFromDetailsRepo(studentattendance)
}
