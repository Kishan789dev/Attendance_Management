package students

import (
	"errors"
	"fmt"
	"time"

	"github.com/kk/attendance_management/bean"
	"github.com/kk/attendance_management/dataBase"
)

// utils:::::::::

type StudentRepository interface {
	GetSid(email string) (error, int)
	AddStudentRepo(student *bean.Student) error
	StudentEntryPunchinRepo(email string) (error, int, int)
	StudentEntryPunchinEntryRepo(sid int) (int, error)
	StudentEntryPunchinEntryTableRepo(aid int) error
	StudentEntryOnlyPunchinRepo(aid int) (int, int, error)
	StudentPunchingRepo(aid int) error
	StudentEntryPunchOutRepo(email string) (error, int, int, int)
	StudentEntryOnlyPunchOutRepo(aid int) (int, int, error)
	StudentPunchingOutRepo(aid int) error
	FetchAttendanceFromDetailsRepo(studentattendance *bean.StudentAttendance) (*[]bean.StudentAttendancetemp, error)
}

type StudentRepositoryImpl struct {
	database dataBase.DataBase
}

func NewStudentRepository(database dataBase.DataBase) *StudentRepositoryImpl {
	return &StudentRepositoryImpl{
		database: database,
	}
}

// utils:::::::::
func (impl *StudentRepositoryImpl) GetSid(email string) (error, int) {

	//
	var sid int
	var student bean.Student
	err := impl.database.Connect().Model(&student).Column("sid").Where("email=?", email).Select(&sid)
	if err != nil {
		return fmt.Errorf("error while getting sid from email reason:%s", err), 0
	}
	return nil, sid

}

func (impl *StudentRepositoryImpl) AddStudentRepo(student *bean.Student) error {

	_, err := impl.database.Connect().Model(student).Insert()
	if err != nil {
		return err
	}

	return nil

}

// punchin

func (impl *StudentRepositoryImpl) StudentEntryPunchinRepo(email string) (error, int, int) {

	var sid int
	var studentattendance bean.StudentAttendance
	var student bean.Student
	err := impl.database.Connect().Model(&student).Column("sid").Where("email=?", email).Select(&sid)
	if err != nil {
		return err, 0, 0
	}

	err = impl.database.Connect().Model(&studentattendance).Where("sid=? and date=? and month=? and year=? ", sid, time.Now().Day(), int(time.Now().Month()), time.Now().Year()).Select() // add date in where claise
	if err != nil {

		return err, 1, sid
	}

	return nil, 2, studentattendance.Aid

}

func (impl *StudentRepositoryImpl) StudentEntryPunchinEntryRepo(sid int) (int, error) {

	studentattendance := bean.StudentAttendance{Sid: sid}
	studentattendance.Date = time.Now().Day()
	studentattendance.Month = int(time.Now().Month())
	studentattendance.Year = time.Now().Year()

	_, err := impl.database.Connect().Model(&studentattendance).Insert()
	if err != nil {
		return 0, err
	}
	return studentattendance.Aid, nil
}

func (impl *StudentRepositoryImpl) StudentEntryPunchinEntryTableRepo(aid int) error {

	punchin := &bean.StudentLogPunchs{
		Aid:  aid,
		Time: time.Now().Add(time.Hour*5 + time.Minute*30),
		Type: 1,
	}

	_, err := impl.database.Connect().Model(punchin).Insert()
	if err != nil {
		return err
	}

	return nil
}

func (impl *StudentRepositoryImpl) StudentEntryOnlyPunchinRepo(aid int) (int, int, error) {

	punchtable := bean.StudentLogPunchs{}

	cnt1, err1 := impl.database.Connect().Model(&punchtable).Where("aid=? and type=?", aid, 1).Count()

	cnt2, err2 := impl.database.Connect().Model(&punchtable).Where("aid=? and type=?", aid, 2).Count()

	if err1 != nil {
		return 0, 0, err1
	}
	if err2 != nil {
		return 0, 0, err2
	}
	return cnt1, cnt2, nil
}

func (impl *StudentRepositoryImpl) StudentPunchingRepo(aid int) error {
	punchtable := &bean.StudentLogPunchs{}

	punchtable.Time = time.Now().Add(time.Hour*5 + time.Minute*30)
	punchtable.Aid = aid
	punchtable.Type = 1
	_, err := impl.database.Connect().Model(punchtable).Insert()

	if err != nil {
		return errors.New(fmt.Sprintf("errored while punching in the student, reason:%s", err.Error()))
	}
	return nil
}

// *******punchout

func (impl *StudentRepositoryImpl) StudentEntryPunchOutRepo(email string) (error, int, int, int) {

	var sid int
	var studentattendance bean.StudentAttendance
	var student bean.Student
	err := impl.database.Connect().Model(&student).Column("sid").Where("email=?", email).Select(&sid)
	if err != nil {
		return err, 0, 0, 0
	}
	studentattendance.Sid = sid

	err = impl.database.Connect().Model(&studentattendance).Where("sid=? and date=? and month=? and year=? ", sid, time.Now().Day(), int(time.Now().Month()), time.Now().Year()).Select() // add date in where claise
	if err != nil {

		return err, 1, sid, 0
	}

	return nil, 2, sid, studentattendance.Aid

}

func (impl *StudentRepositoryImpl) StudentEntryOnlyPunchOutRepo(aid int) (int, int, error) {

	punchtable := bean.StudentLogPunchs{Aid: aid}

	cnt1, err1 := impl.database.Connect().Model(&punchtable).Where("aid=? and type=?", aid, 1).Count()

	cnt2, err2 := impl.database.Connect().Model(&punchtable).Where("aid=? and type=?", aid, 2).Count()

	if err1 != nil {
		return 0, 0, err1
	}
	if err2 != nil {
		return 0, 0, err2
	}
	return cnt1, cnt2, nil
}

func (impl *StudentRepositoryImpl) StudentPunchingOutRepo(aid int) error {

	punchtable := &bean.StudentLogPunchs{Aid: aid}

	punchtable.Time = time.Now().Add(time.Hour*5 + time.Minute*30)
	punchtable.Aid = aid
	punchtable.Type = 1
	_, err := impl.database.Connect().Model(punchtable).Insert()

	if err != nil {
		return errors.New(fmt.Sprintf("errored while punching in the student, reason:%s", err.Error()))
	}
	return nil
}

// Getstudents attendance

func (impl *StudentRepositoryImpl) FetchAttendanceFromDetailsRepo(studentattendance *bean.StudentAttendance) (*[]bean.StudentAttendancetemp, error) {

	var studentattendancedetail []bean.StudentAttendancetemp
	err := impl.database.Connect().Model(&studentattendancedetail).
		ColumnExpr(" DISTINCT student_attendances.date").
		Column("student_attendances.month").
		Column("student_attendances.year").
		Column("student_log_punchs.time").
		Column("student_log_punchs.type").
		Join("inner join student_attendances on student_attendances.aid=student_log_punchs.aid").
		Table("student_log_punchs").
		Where("student_attendances.sid=? AND student_attendances.month=? AND student_attendances.year=?", studentattendance.Sid, studentattendance.Month, studentattendance.Year).
		Select()
	if err != nil {
		return nil, fmt.Errorf("error while getting attendace details of student:%s", err)
	}
	return &studentattendancedetail, nil
}
