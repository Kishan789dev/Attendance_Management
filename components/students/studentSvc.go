package students

import (
	"fmt"
	"log"

	"github.com/kk/attendance_management/bean"
	// repository "github.com/kk/attendance_management/repository"
)

func GetStudentsSvc() ([]bean.Student, error) {

	students, err := GetStudentsRepo()
	return students, err
}

func GetStudentSvc(id int) (bean.Student, error) {

	students, err := GetStudentRepo(id)
	if err != nil {
		return students, err
	}

	return students, nil
}

func AddStudentService(student *bean.Student) error {

	err := AddStudentRepo(student)
	if err != nil {

		return err
	}

	return nil
}

// punchin

func StudentAttendanceWithPunchData(sid int) (error, int) {
	aid, err := StudentEntryPunchinEntryRepo(sid)
	if err != nil {
		return err, 0
	}

	err = StudentEntryPunchinEntryTableRepo(aid)
	if err != nil {
		return err, 1
	}

	return nil, aid
}

func StudentPunchEntryInTable(aid int) (error, string) {
	pi_count, po_count, err := StudentEntryOnlyPunchinSvc(aid)
	if err != nil {
		return err, ""
	}

	if pi_count <= po_count {

		err = StudentPunchingSvc(aid)
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
	err, typ, sid := StudentEntryPunchinRepo(email)
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
	err, typ, sid, aid := StudentEntryPunchOutRepo(email)
	if err != nil {

		return err, typ, sid, aid
	}

	return nil, 3, sid, aid

}

// func StudentAttendanceWithPunchOutData() (error, int) {
// 	// aid, err := StudentEntryPunchinEntryRepo()
// 	// if err != nil {
// 	// 	return err, 0
// 	// }

// 	err := StudentEntryPunchinEntryTableRepo(aid)
// 	if err != nil {
// 		return err, 0
// 	}

// 	return nil, aid
// }

func StudentPunchOutEntryInTable(aid int) (error, string) {
	pi_count, po_count, err := StudentEntryOnlyPunchOutSvc(aid)
	if err != nil {
		return err, ""
	}

	if pi_count > po_count {

		err = StudentPunchingOutRepo(aid)
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
	err, sid := GetSid(email)
	if err != nil {
		return err, 0

	}
	return nil, sid

}
func FetchAttendanceFromDetailsSvc(studentattendance *bean.StudentAttendance) (*[]bean.StudentAttendancetemp, error) {
	return FetchAttendanceFromDetailsRepo(studentattendance)
}
