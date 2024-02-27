package students

import (
	"github.com/kk/attendance_management/bean"
	"github.com/kk/attendance_management/components/users"
)

type StudentService interface {
	AddStudentService(student *bean.Student, userdetails *bean.Userdetails) error
	StudentAttendanceWithPunchData(sid int) (error, int)
	StudentPunchEntryInTable(aid int) (error, string)
	StudentEntryPunchinSvc(email string) (error, int, int)
	StudentEntryPunchOutSvc(email string) (error, int, int, int)
	StudentPunchOutEntryInTable(aid int) (error, string)
	GetStudentsAttendanceEmailSvc(email string) (error, int)
	FetchAttendanceFromDetailsSvc(studentattendance *bean.StudentAttendance) (*[]bean.StudentAttendancetemp, error)
}

type StudentServiceImpl struct {
	studentRepo StudentRepository
	userrest    users.UserRest
}

func NewStudentservice(studentRepo StudentRepository, userrest users.UserRest) *StudentServiceImpl {
	return &StudentServiceImpl{
		studentRepo: studentRepo,
		userrest:    userrest,
	}
}

func (impl *StudentServiceImpl) AddStudentService(student *bean.Student, userdetails *bean.Userdetails) error {

	err := impl.studentRepo.AddStudentRepo(student)
	if err == nil {
		impl.userrest.AddUser(userdetails.Email, 1, userdetails.Password)
		return err
	}

	return nil
}

func (impl *StudentServiceImpl) StudentAttendanceWithPunchData(sid int) (error, int) {
	aid, err := impl.studentRepo.StudentEntryPunchinEntryRepo(sid)
	if err != nil {
		return err, 0
	}

	err = impl.studentRepo.StudentEntryPunchinEntryTableRepo(aid)
	if err != nil {
		return err, 1
	}

	return nil, aid
}

func (impl *StudentServiceImpl) StudentPunchEntryInTable(aid int) (error, string) {
	pi_count, po_count, err := impl.studentRepo.StudentEntryOnlyPunchinRepo(aid)
	if err != nil {
		return err, ""
	}

	if pi_count <= po_count {

		err = impl.studentRepo.StudentPunchingRepo(aid)
		if err != nil {

			return err, ""
		}
		return err, "punch in successful"

	} else {
		return err, "already punch in"

	}

}

func (impl *StudentServiceImpl) StudentEntryPunchinSvc(email string) (error, int, int) {
	err, typ, sid := impl.studentRepo.StudentEntryPunchinRepo(email)
	if err != nil {
		return err, typ, sid
	}

	return nil, typ, sid

}

func (impl *StudentServiceImpl) StudentEntryPunchOutSvc(email string) (error, int, int, int) {
	err, typ, sid, aid := impl.studentRepo.StudentEntryPunchOutRepo(email)
	if err != nil {

		return err, typ, sid, aid
	}

	return nil, 3, sid, aid

}

func (impl *StudentServiceImpl) StudentPunchOutEntryInTable(aid int) (error, string) {
	pi_count, po_count, err := impl.studentRepo.StudentEntryOnlyPunchOutRepo(aid)
	if err != nil {
		return err, ""
	}

	if pi_count > po_count {

		err = impl.studentRepo.StudentPunchingOutRepo(aid)
		if err != nil {

			return err, ""
		}
		return nil, "Puchout successfully"

	} else {

		return nil, "you have already punch out"

	}

}

func (impl *StudentServiceImpl) GetStudentsAttendanceEmailSvc(email string) (error, int) {
	err, sid := impl.studentRepo.GetSid(email)
	if err != nil {
		return err, 0

	}
	return nil, sid

}
func (impl *StudentServiceImpl) FetchAttendanceFromDetailsSvc(studentattendance *bean.StudentAttendance) (*[]bean.StudentAttendancetemp, error) {
	return impl.studentRepo.FetchAttendanceFromDetailsRepo(studentattendance)
}
