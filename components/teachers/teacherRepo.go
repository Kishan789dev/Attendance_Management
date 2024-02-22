package teachers

import (
	"errors"
	"fmt"
	"log"
	"time"

	// "github.com/go-pg/pg"
	"github.com/go-pg/pg"
	"github.com/kk/attendance_management/bean"
	"github.com/kk/attendance_management/dataBase"
)

func AddTeacherRepo(teacher *bean.Teacher) error {
	db := dataBase.Connect()
	defer db.Close()
	fmt.Println("teacherrrr", teacher)
	fmt.Println("teacherrrr", teacher)

	_, err := db.Model(teacher).Insert()
	if err != nil {
		return fmt.Errorf("error during inserting teacher data in teacher table ,reason:%s", err)
	}
	return nil

}

// punchin

func TeacherEntryPunchinRepo(email string) (error, int, int) {
	db := dataBase.Connect()
	defer db.Close()
	var tid int
	var teacherattendance bean.TeacherAttendance
	var teacher bean.Teacher
	err := db.Model(&teacher).Column("tid").Where("email=?", email).Select(&tid)
	fmt.Println("repo tid", tid)
	if err != nil {
		log.Println("err1")
		return err, 0, 0
	}
	log.Println("pppppppppppppppppp", teacherattendance)
	// teacherattendance.Tid = tid
	log.Println("sssstid", tid)

	err = db.Model(&teacherattendance).Where("tid=? and date=? and month=? and year=? ", tid, time.Now().Day(), int(time.Now().Month()), time.Now().Year()).Select() // add date in where claise
	if err != nil {
		log.Println("mmmm", teacherattendance)

		return err, 1, tid
	}
	fmt.Println(teacherattendance)

	log.Println("lllll", teacherattendance)
	return nil, 2, teacherattendance.Aid

}

func TeacherEntryPunchinEntryRepo(tid int) (int, error) {
	fmt.Println("errror", tid)

	teacherattendance := bean.TeacherAttendance{Tid: tid}
	teacherattendance.Date = time.Now().Day()
	teacherattendance.Month = int(time.Now().Month())
	teacherattendance.Year = time.Now().Year()
	db := dataBase.Connect()
	defer db.Close()
	fmt.Println("sa", teacherattendance)
	_, err := db.Model(&teacherattendance).Insert()
	if err != nil {
		fmt.Println("errror", err)
		return 0, err
	}

	return teacherattendance.Aid, nil
}

func TeacherEntryPunchinEntryTableRepo(aid int) error {
	db := dataBase.Connect()
	defer db.Close()

	punchin := &bean.TeacherLogPunchs{
		Aid:  aid,
		Time: time.Now().Add(time.Hour*5 + time.Minute*30),
		Type: 1,
	}
	log.Println("aiiiiiiiiiiiid", aid)
	log.Println("punchin", punchin)

	_, err := db.Model(punchin).Insert()
	if err != nil {
		return err
	}
	log.Println("liiiiiiiiiiiid", punchin)

	return nil
}

func TeacherEntryOnlyPunchinSvc(aid int) (int, int, error) {
	db := dataBase.Connect()
	defer db.Close()
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

func TeacherPunchingSvc(aid int) error {
	log.Println("raunitttttttttttttttttt", aid)
	punchtable := &bean.TeacherLogPunchs{}
	db := dataBase.Connect()
	defer db.Close()
	punchtable.Time = time.Now().Add(time.Hour*5 + time.Minute*30)
	punchtable.Aid = aid
	punchtable.Type = 1
	_, err := db.Model(punchtable).Insert()

	if err != nil {
		return errors.New(fmt.Sprintf("errored while punching in the teacher, reason:%s", err.Error()))
	}
	return nil
}

// *******punchout

func TeacherEntryPunchOutRepo(email string) (error, int, int, int) {
	db := dataBase.Connect()
	defer db.Close()
	var tid int
	var teacherattendance bean.TeacherAttendance
	var teacher bean.Teacher
	err := db.Model(&teacher).Column("tid").Where("email=?", email).Select(&tid)
	fmt.Println("repo tid", tid)
	if err != nil {
		log.Println("err1")
		return err, 0, 0, 0
	}
	teacherattendance.Tid = tid
	log.Println("teacheratttendance", teacherattendance)

	err = db.Model(&teacherattendance).Where("tid=? and date=? and month=? and year=? ", tid, time.Now().Day(), int(time.Now().Month()), time.Now().Year()).Select() // add date in where claise
	if err != nil {
		log.Println("teacheratttendance", teacherattendance)

		log.Println("lllll", teacherattendance)
		return err, 1, tid, 0
	}

	return nil, 2, tid, teacherattendance.Aid

}

func TeacherEntryOnlyPunchOutSvc(aid int) (int, int, error) {
	db := dataBase.Connect()
	defer db.Close()
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

func TeacherPunchingOutRepo(aid int) error {

	punchtable := &bean.TeacherLogPunchs{Aid: aid}
	db := dataBase.Connect()
	defer db.Close()
	punchtable.Time = time.Now().Add(time.Hour*5 + time.Minute*30)
	punchtable.Aid = aid
	punchtable.Type = 2
	_, err := db.Model(punchtable).Insert()

	if err != nil {
		return errors.New(fmt.Sprintf("errored while punching in the teacher, reason:%s", err.Error()))
	}
	return nil
}

// Get attendance of teacher

func GetTeacherIDByEmail(email string) (int, error) {
	db := dataBase.Connect()
	defer db.Close()

	var tid int
	var teacher bean.Teacher
	err := db.Model(&teacher).Column("tid").Where("email=?", email).Select(&tid)

	return tid, err
}

// GetTeacherAttendance

func GetTeacherattendanceRepoTidGetting(email string, tid *int) error {
	db := dataBase.Connect()
	defer db.Close()

	err := db.Model(&bean.Teacher{}).Column("tid").Where("email=?", email).Select(tid)

	if err != nil {
		if err == pg.ErrNoRows {
			return fmt.Errorf("teacher with details don't exist")
		}
		return fmt.Errorf("error in data fetching:%s", err)

	}
	return nil

}

func GetTeacherAttendanceDetailsRepo(tid int, month int, year int) (error, []bean.TeacherAttendancetemp) {
	db := dataBase.Connect()
	defer db.Close()

	var teacherattendancedetail []bean.TeacherAttendancetemp
	err := db.Model(&teacherattendancedetail).
		// ColumnExpr(" DISTINCT teacher_attendances.date").
		// Column("teacher_attendances.month").
		// Column("teacher_attendances.year").
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
		fmt.Println("kokokokokokokokokokok")

		return fmt.Errorf("error in data fetching:%s", err), teacherattendancedetail

	}
	fmt.Println("eeeeeeeeeeeeee")
	return nil, teacherattendancedetail

}

// Get class attendance

func GetClassattendanceRepo(classtemp *bean.Classtemp) (error, *[]bean.ClasstempRes) {
	db := dataBase.Connect()
	defer db.Close()

	var classdata []bean.ClasstempRes

	// err := db.Model(&student).Where("class =?", classtemp.Class).Select()
	err := db.Model(&classdata).
		// ColumnExpr(" DISTINCT students.sid").
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
		// if err == pg.ErrNoRows {
		// 	return fmt.Errorf("ateendance with details don't exist"), teacherattendancedetail
		// }
		return fmt.Errorf("error in data fetching:%s", err), &classdata

	}
	return nil, &classdata

}
