package teachers

import (
	"github.com/kk/attendance_management/bean"
	"github.com/kk/attendance_management/components/users"
)

type TeacherSvc interface {
	AddTeacherSvc(teacher *bean.Teacher, userdetails *bean.Userdetails) error
	TeacherAttendanceWithPunchData(tid int) (error, int)
	TeacherPunchEntryInTable(aid int) (error, string)
	TeacherEntryPunchinSvc(email string) (error, int, int)
	TeacherEntryPunchOutSvc(email string) (error, int, int, int)
	TeacherPunchOutEntryInTable(aid int) (error, string)
	GetTeacherattendanceSvcTidGetting(email string, tid *int) error
	GetTeacherAttendanceDetailsSvc(tid int, month int, year int) (error, []bean.TeacherAttendancetemp)
	GetClassattendanceSvc(classtemp *bean.Classtemp) (error, *[]bean.ClasstempRes)
}

type TeacherSvcImpl struct {
	teacherrepo TeacherRepo
	userrest    users.UserRest
}

func NewTeacherSvc(teacherrepo TeacherRepo, userrest users.UserRest) *TeacherSvcImpl {
	return &TeacherSvcImpl{
		teacherrepo: teacherrepo,
		userrest:    userrest,
	}
}

func (impl *TeacherSvcImpl) AddTeacherSvc(teacher *bean.Teacher, userdetails *bean.Userdetails) error {
	err := impl.teacherrepo.AddTeacherRepo(teacher)
	if err == nil {
		impl.userrest.AddUser(userdetails.Email, 2, userdetails.Password)
		return nil
	}
	return err

}

func (impl *TeacherSvcImpl) TeacherAttendanceWithPunchData(tid int) (error, int) {
	aid, err := impl.teacherrepo.TeacherEntryPunchinEntryRepo(tid)
	if err != nil {
		return err, 0
	}

	err = impl.teacherrepo.TeacherEntryPunchinEntryTableRepo(aid)
	if err != nil {
		return err, 1
	}

	return nil, aid
}

func (impl *TeacherSvcImpl) TeacherPunchEntryInTable(aid int) (error, string) {
	pi_count, po_count, err := impl.teacherrepo.TeacherEntryOnlyPunchinRepo(aid)
	if err != nil {
		return err, ""
	}

	if pi_count <= po_count {

		err = impl.teacherrepo.TeacherPunchingRepo(aid)
		if err != nil {

			return err, ""
		}
		return err, "punch in successful"

	} else {
		return err, "already punch in"

	}

}

func (impl *TeacherSvcImpl) TeacherEntryPunchinSvc(email string) (error, int, int) {
	err, typ, tid := impl.teacherrepo.TeacherEntryPunchinRepo(email)
	if err != nil {
		return err, typ, tid
	}

	return nil, typ, tid

}

func (impl *TeacherSvcImpl) TeacherEntryPunchOutSvc(email string) (error, int, int, int) {
	err, typ, tid, aid := impl.teacherrepo.TeacherEntryPunchOutRepo(email)
	if err != nil {

		return err, typ, tid, aid
	}

	return nil, 3, tid, aid

}

func (impl *TeacherSvcImpl) TeacherPunchOutEntryInTable(aid int) (error, string) {
	pi_count, po_count, err := impl.teacherrepo.TeacherEntryOnlyPunchOutRepo(aid)
	if err != nil {
		return err, ""
	}

	if pi_count > po_count {

		err = impl.teacherrepo.TeacherPunchingOutRepo(aid)
		if err != nil {

			return err, ""
		}
		return nil, "Puchout successfully"

	} else {

		return nil, "you have already punch out"

	}

}

func (impl *TeacherSvcImpl) GetTeacherattendanceSvcTidGetting(email string, tid *int) error {

	return impl.teacherrepo.GetTeacherattendanceRepoTidGetting(email, tid)

}

func (impl *TeacherSvcImpl) GetTeacherAttendanceDetailsSvc(tid int, month int, year int) (error, []bean.TeacherAttendancetemp) {

	return impl.teacherrepo.GetTeacherAttendanceDetailsRepo(tid, month, year)

}

func (impl *TeacherSvcImpl) GetClassattendanceSvc(classtemp *bean.Classtemp) (error, *[]bean.ClasstempRes) {

	return impl.teacherrepo.GetClassattendanceRepo(classtemp)
}
