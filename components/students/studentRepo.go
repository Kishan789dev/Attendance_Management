package students

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/kk/attendance_management/bean"
	"github.com/kk/attendance_management/dataBase"
)

// utils:::::::::
func GetSid(email string) (error, int) {
	db := dataBase.Connect()
	defer db.Close()
	//
	var sid int
	var student bean.Student
	err := db.Model(&student).Column("sid").Where("email=?", email).Select(&sid)
	if err != nil {
		return fmt.Errorf("error while getting sid from email reason:%s", err), 0
	}
	return nil, sid

}

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

func AddStudentService(student *bean.Student) error {
	db := dataBase.Connect()
	defer db.Close()

	_, err := db.Model(student).Insert()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil

}

// punchin

func StudentEntryPunchinRepo(email string) (error, int, int) {
	db := dataBase.Connect()
	defer db.Close()
	var sid int
	var studentattendance bean.StudentAttendance
	var student bean.Student
	err := db.Model(&student).Column("sid").Where("email=?", email).Select(&sid)
	fmt.Println("repo sid", sid)
	if err != nil {
		log.Println("err1")
		return err, 0, 0
	}
	log.Println("pppppppppppppppppp", studentattendance)
	// studentattendance.Sid = sid
	log.Println("sssssid", sid)

	err = db.Model(&studentattendance).Where("sid=? and date=? and month=? and year=? ", sid, time.Now().Day(), int(time.Now().Month()), time.Now().Year()).Select() // add date in where claise
	if err != nil {
		log.Println("mmmm", studentattendance)

		return err, 1, sid
	}
	fmt.Println(studentattendance)

	log.Println("lllll", studentattendance)
	return nil, 2, studentattendance.Aid

}

func StudentEntryPunchinEntryRepo(sid int) (int, error) {
	fmt.Println("errror", sid)

	studentattendance := bean.StudentAttendance{Sid: sid}
	studentattendance.Date = time.Now().Day()
	studentattendance.Month = int(time.Now().Month())
	studentattendance.Year = time.Now().Year()
	db := dataBase.Connect()
	defer db.Close()
	fmt.Println("sa", studentattendance)
	_, err := db.Model(&studentattendance).Insert()
	if err != nil {
		fmt.Println("errror", err)
		return 0, err
	}
	fmt.Println(" Raunit verma ", studentattendance.Aid)
	return studentattendance.Aid, nil
}

func StudentEntryPunchinEntryTableRepo(aid int) error {
	db := dataBase.Connect()
	defer db.Close()

	punchin := &bean.StudentLogPunchs{
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

func StudentEntryOnlyPunchinSvc(aid int) (int, int, error) {
	db := dataBase.Connect()
	defer db.Close()
	punchtable := bean.StudentLogPunchs{}

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

func StudentPunchingSvc(aid int) error {
	log.Println("raunitttttttttttttttttt", aid)
	punchtable := &bean.StudentLogPunchs{}
	db := dataBase.Connect()
	defer db.Close()
	punchtable.Time = time.Now().Add(time.Hour*5 + time.Minute*30)
	punchtable.Aid = aid
	punchtable.Type = 1
	_, err := db.Model(punchtable).Insert()

	if err != nil {
		log.Println("ferrooooooooooooooo", err)
		return errors.New(fmt.Sprintf("errored while punching in the student, reason:%s", err.Error()))
	}
	return nil
}

// *******punchout

func StudentEntryPunchOutRepo(email string) (error, int, int, int) {
	db := dataBase.Connect()
	defer db.Close()
	var sid int
	var studentattendance bean.StudentAttendance
	var student bean.Student
	err := db.Model(&student).Column("sid").Where("email=?", email).Select(&sid)
	fmt.Println("repo sid", sid)
	if err != nil {
		log.Println("err1")
		return err, 0, 0, 0
	}
	studentattendance.Sid = sid
	log.Println("studentatttendance", studentattendance)

	err = db.Model(&studentattendance).Where("sid=? and date=? and month=? and year=? ", sid, time.Now().Day(), int(time.Now().Month()), time.Now().Year()).Select() // add date in where claise
	if err != nil {
		log.Println("studentatttendance", studentattendance)

		log.Println("lllll", studentattendance)
		return err, 1, sid, 0
	}

	return nil, 2, sid, studentattendance.Aid

}

func StudentEntryOnlyPunchOutSvc(aid int) (int, int, error) {
	db := dataBase.Connect()
	defer db.Close()
	punchtable := bean.StudentLogPunchs{Aid: aid}

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

func StudentPunchingOutRepo(aid int) error {

	punchtable := &bean.StudentLogPunchs{Aid: aid}
	db := dataBase.Connect()
	defer db.Close()
	punchtable.Time = time.Now().Add(time.Hour*5 + time.Minute*30)
	punchtable.Aid = aid
	punchtable.Type = 2
	_, err := db.Model(punchtable).Insert()

	if err != nil {
		return errors.New(fmt.Sprintf("errored while punching in the student, reason:%s", err.Error()))
	}
	return nil
}

// Getstudents attendance

func FetchAttendanceFromDetailsRepo(studentattendance *bean.StudentAttendance) (*[]bean.StudentAttendancetemp, error) {
	db := dataBase.Connect()
	defer db.Close()
	var studentattendancedetail []bean.StudentAttendancetemp
	err := db.Model(&studentattendancedetail).
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
		// log.Println("fgfkjfgjkfgjkgfkjfgjkhfgkjhjkghgjkhgfjkhgfkjfgkjjgjkfjfgjkhfkjgfkjhfgjkhjjfgjkfgkjjfgjkhjgfjkjgfkj")
		return nil, fmt.Errorf("error while getting attendace details of student:%s", err)
	}
	return &studentattendancedetail, nil
}
