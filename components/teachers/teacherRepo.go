package teachers

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-pg/pg"
	"github.com/kk/attendance_management/bean"
	"github.com/kk/attendance_management/dataBase"
)

type TeacherRepo interface {
	AddTeacherRepo(teacher *bean.Teacher) error
	TeacherEntryPunchinRepo(email string) (error, int, int)
	TeacherEntryPunchinEntryRepo(tid int) (int, error)
	TeacherEntryPunchinEntryTableRepo(aid int) error

	TeacherEntryOnlyPunchinRepo(aid int) (int, int, error)
	TeacherPunchingRepo(aid int) error
	TeacherEntryPunchOutRepo(email string) (error, int, int, int)
	TeacherEntryOnlyPunchOutRepo(aid int) (int, int, error)
	TeacherPunchingOutRepo(aid int) error
	GetTeacherIDByEmail(email string) (int, error)
	GetTeacherattendanceRepoTidGetting(email string, tid *int) error
	GetTeacherAttendanceDetailsRepo(tid int, month int, year int) (error, []bean.TeacherAttendancetemp)
	GetClassattendanceRepo(classtemp *bean.Classtemp) (error, *[]bean.ClasstempRes)
}

type TeacherRepositoryImpl struct {
	database dataBase.DataBase
}

func NewTeacherRepositoryImpl(database dataBase.DataBase) *TeacherRepositoryImpl {
	return &TeacherRepositoryImpl{
		database: database,
	}
}

func (impl *TeacherRepositoryImpl) AddTeacherRepo(teacher *bean.Teacher) error {
	db := impl.database.Connect()

	_, err := db.Model(teacher).Insert()
	if err != nil {
		return fmt.Errorf("error during inserting teacher data in teacher table ,reason:%s", err)
	}
	return nil

}
func (impl *TeacherRepositoryImpl) TeacherEntryPunchinRepo(email string) (error, int, int) {
	db := impl.database.Connect()

	var tid int
	var teacherattendance bean.TeacherAttendance
	var teacher bean.Teacher
	err := db.Model(&teacher).Column("tid").Where("email=?", email).Select(&tid)
	if err != nil {
		return err, 0, 0
	}
	err = impl.database.Connect().Model(&teacherattendance).Where("tid=? and date=? and month=? and year=? ", tid, time.Now().Day(), int(time.Now().Month()), time.Now().Year()).Select() // add date in where claise
	if err != nil {

		return err, 1, tid
	}

	return nil, 2, teacherattendance.Aid

}
func (impl *TeacherRepositoryImpl) TeacherEntryPunchinEntryRepo(tid int) (int, error) {

	teacherattendance := bean.TeacherAttendance{Tid: tid}
	teacherattendance.Date = time.Now().Day()
	teacherattendance.Month = int(time.Now().Month())
	teacherattendance.Year = time.Now().Year()
	db := impl.database.Connect()

	_, err := db.Model(&teacherattendance).Insert()
	if err != nil {
		return 0, err
	}

	return teacherattendance.Aid, nil
}
func (impl *TeacherRepositoryImpl) TeacherEntryPunchinEntryTableRepo(aid int) error {
	db := impl.database.Connect()

	punchin := &bean.TeacherLogPunchs{
		Aid:  aid,
		Time: time.Now().Add(time.Hour*5 + time.Minute*30),
		Type: 1,
	}

	_, err := db.Model(punchin).Insert()
	if err != nil {
		return err
	}

	return nil
}
func (impl *TeacherRepositoryImpl) TeacherEntryOnlyPunchinRepo(aid int) (int, int, error) {
	db := impl.database.Connect()

	punchtable := bean.TeacherLogPunchs{}

	cnt1, err1 := db.Model(&punchtable).Where("aid=? and type=?", aid, 1).Count()

	cnt2, err2 := db.Model(&punchtable).Where("aid=? and type=?", aid, 2).Count()

	if err1 != nil {
		return 0, 0, err1
	}
	if err2 != nil {
		return 0, 0, err2
	}
	return cnt1, cnt2, nil
}
func (impl *TeacherRepositoryImpl) TeacherPunchingRepo(aid int) error {
	punchtable := &bean.TeacherLogPunchs{}
	db := impl.database.Connect()

	punchtable.Time = time.Now().Add(time.Hour*5 + time.Minute*30)
	punchtable.Aid = aid
	punchtable.Type = 1
	_, err := db.Model(punchtable).Insert()

	if err != nil {
		return errors.New(fmt.Sprintf("errored while punching in the teacher, reason:%s", err.Error()))
	}
	return nil
}
func (impl *TeacherRepositoryImpl) TeacherEntryPunchOutRepo(email string) (error, int, int, int) {
	db := impl.database.Connect()

	var tid int
	var teacherattendance bean.TeacherAttendance
	var teacher bean.Teacher
	err := db.Model(&teacher).Column("tid").Where("email=?", email).Select(&tid)
	if err != nil {
		return err, 0, 0, 0
	}
	teacherattendance.Tid = tid
	err = db.Model(&teacherattendance).Where("tid=? and date=? and month=? and year=? ", tid, time.Now().Day(), int(time.Now().Month()), time.Now().Year()).Select() // add date in where claise

	if err != nil {

		return err, 1, tid, 0
	}

	return nil, 2, tid, teacherattendance.Aid

}
func (impl *TeacherRepositoryImpl) TeacherEntryOnlyPunchOutRepo(aid int) (int, int, error) {
	db := impl.database.Connect()

	punchtable := bean.TeacherLogPunchs{Aid: aid}

	cnt1, err1 := db.Model(&punchtable).Where("aid=? and type=?", aid, 1).Count()

	cnt2, err2 := db.Model(&punchtable).Where("aid=? and type=?", aid, 2).Count()

	if err1 != nil {
		return 0, 0, err1
	}
	if err2 != nil {
		return 0, 0, err2
	}
	return cnt1, cnt2, nil
}
func (impl *TeacherRepositoryImpl) TeacherPunchingOutRepo(aid int) error {

	db := impl.database.Connect()
	punchtable := &bean.TeacherLogPunchs{Aid: aid}

	punchtable.Time = time.Now().Add(time.Hour*5 + time.Minute*30)
	punchtable.Aid = aid
	punchtable.Type = 2
	_, err := db.Model(punchtable).Insert()

	if err != nil {
		return errors.New(fmt.Sprintf("errored while punching in the teacher, reason:%s", err.Error()))
	}
	return nil
}
func (impl *TeacherRepositoryImpl) GetTeacherIDByEmail(email string) (int, error) {
	db := impl.database.Connect()

	var tid int
	var teacher bean.Teacher
	err := db.Model(&teacher).Column("tid").Where("email=?", email).Select(&tid)

	return tid, err
}
func (impl *TeacherRepositoryImpl) GetTeacherattendanceRepoTidGetting(email string, tid *int) error {
	db := impl.database.Connect()

	err := db.Model(&bean.Teacher{}).Column("tid").Where("email=?", email).Select(tid)

	if err != nil {
		if err == pg.ErrNoRows {
			return fmt.Errorf("teacher with details don't exist")
		}
		return fmt.Errorf("error in data fetching:%s", err)

	}
	return nil

}
func (impl *TeacherRepositoryImpl) GetTeacherAttendanceDetailsRepo(tid int, month int, year int) (error, []bean.TeacherAttendancetemp) {
	db := impl.database.Connect()

	var teacherattendancedetail []bean.TeacherAttendancetemp
	err := db.Model(&teacherattendancedetail).
		ColumnExpr(" DISTINCT teacher_log_punchs.time").
		Column("teacher_log_punchs.type").
		Join("inner join teacher_attendances on teacher_attendances.aid=teacher_log_punchs.aid").
		Table("teacher_log_punchs").
		Where("teacher_attendances.tid=? AND teacher_attendances.month=? AND teacher_attendances.year=?", tid, month, year).
		Order("teacher_log_punchs.time ASC").
		Select()
	if len(teacherattendancedetail) == 0 {
		return fmt.Errorf("teacher doesn't exist with details"), teacherattendancedetail
	}

	if err != nil {

		return fmt.Errorf("error in data fetching:%s", err), teacherattendancedetail

	}
	return nil, teacherattendancedetail

}
func (impl *TeacherRepositoryImpl) GetClassattendanceRepo(classtemp *bean.Classtemp) (error, *[]bean.ClasstempRes) {
	db := impl.database.Connect()

	var classdata []bean.ClasstempRes

	err := db.Model(&classdata).
		ColumnExpr(" DISTINCT students.name").
		Join("INNER JOIN student_attendances on student_attendances.sid=students.sid").
		Table("students").
		Where("student_attendances.date=? AND student_attendances.month=? AND student_attendances.year=?", classtemp.Date, classtemp.Month, classtemp.Year).
		Where("students.class =?", classtemp.Class).
		Select()

	if len(classdata) == 0 {
		return fmt.Errorf("invalid data"), &classdata
	}
	if err != nil {
		return fmt.Errorf("error in data fetching:%s", err), &classdata

	}
	return nil, &classdata

}
