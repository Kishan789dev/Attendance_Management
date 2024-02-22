package teachers

import (
	"fmt"
	"log"

	auth "github.com/kk/attendance_management/authentication"
	"github.com/kk/attendance_management/bean"
	repository "github.com/kk/attendance_management/repository"
)

func AddTeacherSvc(teacher *bean.Teacher, userdetails *bean.Userdetails) error {

	err := repository.AddTeacherRepo(teacher)
	if err == nil {
		auth.AddUser(userdetails.Email, 2, userdetails.Password)
		return nil
	}
	return err

}

// punchin

func TeacherAttendanceWithPunchData(tid int) (error, int) {
	aid, err := repository.TeacherEntryPunchinEntryRepo(tid)
	if err != nil {
		return err, 0
	}

	err = repository.TeacherEntryPunchinEntryTableRepo(aid)
	if err != nil {
		return err, 1
	}

	return nil, aid
}

func TeacherPunchEntryInTable(aid int) (error, string) {
	pi_count, po_count, err := repository.TeacherEntryOnlyPunchinSvc(aid)
	if err != nil {
		return err, ""
	}

	if pi_count <= po_count {

		err = repository.TeacherPunchingSvc(aid)
		if err != nil {

			return err, ""
		}
		return err, "punch in successful"

	} else {
		return err, "already punch in"

	}

}

func TeacherEntryPunchinSvc(email string) (error, int, int) {
	err, typ, tid := repository.TeacherEntryPunchinRepo(email)
	fmt.Println("jdsjsfjsjdsjsdjdsjdsjdsjdssdjjs", tid)
	if err != nil {
		log.Println("tidddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd")
		return err, typ, tid
	}

	return nil, typ, tid

}

// *********punchout
// var aid int

func TeacherEntryPunchOutSvc(email string) (error, int, int, int) {
	err, typ, tid, aid := repository.TeacherEntryPunchOutRepo(email)
	if err != nil {

		return err, typ, tid, aid
	}

	return nil, 3, tid, aid

}

// func TeacherAttendanceWithPunchOutData() (error, int) {
// 	// aid, err := repository.TeacherEntryPunchinEntryRepo()
// 	// if err != nil {
// 	// 	return err, 0
// 	// }

// 	err := repository.TeacherEntryPunchinEntryTableRepo(aid)
// 	if err != nil {
// 		return err, 0
// 	}

// 	return nil, aid
// }

func TeacherPunchOutEntryInTable(aid int) (error, string) {
	pi_count, po_count, err := repository.TeacherEntryOnlyPunchOutSvc(aid)
	if err != nil {
		return err, ""
	}

	if pi_count > po_count {

		err = repository.TeacherPunchingOutRepo(aid)
		if err != nil {

			return err, ""
		}
		return nil, "Puchout successfully"

	} else {

		return nil, "you have already punch out"

	}

}

// GetTeacherAttendance

func GetTeacherattendanceSvcTidGetting(email string, tid *int) error {

	return repository.GetTeacherattendanceRepoTidGetting(email, tid)

}

func GetTeacherAttendanceDetailsSvc(tid int, month int, year int) (error, []bean.TeacherAttendancetemp) {

	return repository.GetTeacherAttendanceDetailsRepo(tid, month, year)

}

// Get class attendance

func GetClassattendanceSvc(classtemp *bean.Classtemp) (error, *[]bean.ClasstempRes) {

	return repository.GetClassattendanceRepo(classtemp)
}
