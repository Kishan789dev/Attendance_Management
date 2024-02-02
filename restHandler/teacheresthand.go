package restHandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-pg/pg"
	"github.com/gorilla/mux"
	bean "github.com/kk/attendance_management/bean"
	"github.com/kk/attendance_management/dataBase"
)

// func GetTeachers(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	var teacher []Teacher

// 	json.NewEncoder(w).Encode(teacher)
// 	return
// }

// func getTeacher(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(r)

// 	var teacher Teacher

// 	json.NewEncoder(w).Encode(teacher)

// }

// func addTeacher(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var teacher Teacher
// 	json.NewDecoder(r.Body).Decode(&teacher)
// 	// id := teacher.TeacherId
// 	// newEmail := teacher.TeacherEmail
// 	json.NewEncoder(w).Encode(teacher)

// }

// func updateTeacher(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(r)

// 	var teacher Teacher

// 	json.NewDecoder(r.Body).Decode(&teacher)
// 	// save

// 	json.NewEncoder(w).Encode(teacher)

// }

// func deleteTeacher(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(r)

// 	var teacher Teacher
// 	// save

// 	json.NewEncoder(w).Encode("deleted successful")

// }

func GetTeachers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := dataBase.Connect()
	defer db.Close()
	var teachers []bean.Teacher
	if err := db.Model(&teachers).Select(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//
	json.NewEncoder(w).Encode(teachers)

}

func GetTeacher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// log.Println("dsjflsdflsdjflksd")
	params := mux.Vars(r)

	db := dataBase.Connect()
	defer db.Close()

	// teacher_id := params["id"]
	teacher_id := params["id"]
	trr, err := strconv.Atoi(teacher_id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// log.Println(err)

	teachers := &bean.Teacher{Tid: trr}

	if err := db.Model(teachers).Where("tid=?", trr).Select(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(teachers)

}

func AddTeacher(w http.ResponseWriter, r *http.Request) {
	fmt.Print("hello2")
	w.Header().Set("Content-Type", "application/json")

	teacher := bean.Teacher{}
	_ = json.NewDecoder(r.Body).Decode(&teacher)

	db := dataBase.Connect()
	defer db.Close()
	// teacher.Id = uuid.New().String()

	if _, err := db.Model(&teacher).Insert(); err != nil {
		log.Println(err)
		// json.NewEncoder(w).Encode("error is line no 77")

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(teacher)

}

func UpdateTeacher(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	db := dataBase.Connect()
	defer db.Close()

	params := mux.Vars(r)

	teacher_id := params["id"]
	trr, err := strconv.Atoi(teacher_id)
	log.Println(err)
	teachers := &bean.Teacher{Tid: trr}

	_ = json.NewDecoder(r.Body).Decode(&teachers)
	yy, err := db.Model(teachers).Where("tid=?", teacher_id).Set("name= ?,address=?,email=?", teachers.Name, teachers.Address, teachers.Email).Update()
	log.Println(yy)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(teachers)

}

func DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	db := dataBase.Connect()
	defer db.Close()

	teacher_id := params["id"]

	trr, err := strconv.Atoi(teacher_id)
	log.Println(err)

	teachers := &bean.Teacher{Tid: trr}
	result, err := db.Model(teachers).Where("tid=?", teacher_id).Delete()

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if result != nil {
		json.NewEncoder(w).Encode("deleted successfully")
	}

}

// *****************TEACHER ATTENDANCE**************************

// func AddTeacherattendace(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var teacherattendace bean.TeacherAttendance
// 	json.NewDecoder(r.Body).Decode(&teacherattendace)
// 	json.NewEncoder(w).Encode(teacherattendace)
// }

// func getTeacherattendace(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(r)
// 	id := params["id"]
// 	month := params["month"]
// 	year := params["year"]

// 	var teacherattendace AttendanceTeacher
// 	json.NewDecoder(r.Body).Decode(&teacherattendace)
// 	json.NewEncoder(w).Encode(teacherattendace)

// }

func TeacherEntryPunchin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("kdjfsdkjsjk")
	teacherattendance := bean.TeacherAttendance{}

	_ = json.NewDecoder(r.Body).Decode(&teacherattendance)

	db := dataBase.Connect()

	defer db.Close()

	// err := db.Model(&teacherattendance).Where("id=? and date=?", teacherattendance.Sid, teacherattendance.Date).Select() // add date in where clause
	err := db.Model(&teacherattendance).Where("Tid=? and date=? and month=? and year=? ", teacherattendance.Tid, time.Now().Day(), int(time.Now().Month()), time.Now().Year()).Select() // add date in where claise

	if err == pg.ErrNoRows {
		//  teacherattendace.PunchIntime=time.Now()
		log.Println(teacherattendance.Tid)

		teacherattendance.Date = time.Now().Day()
		teacherattendance.Month = int(time.Now().Month())
		teacherattendance.Year = time.Now().Year()

		_, err := db.Model(&teacherattendance).Insert()
		if err != nil {
			log.Println("166")
			log.Println(err)

			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// log punch in

		punchin := &bean.TeacherLogPunchs{
			Aid:  teacherattendance.Aid,
			Time: time.Now().Add(time.Hour*5 + time.Minute*30),
			Type: 1,
		}
		_, err = db.Model(punchin).Insert()

		if err != nil {
			log.Println("182")

			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	} else if err != nil {
		log.Println("190")

		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return

	} else {

		aid := teacherattendance.Aid

		punchtable := bean.TeacherLogPunchs{Aid: aid}

		pi_count, err := db.Model(&punchtable).Where("aid=? and type=?", aid, 1).Count()

		if err != nil {
			log.Println("205")

			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		po_count, err := db.Model(&punchtable).Where("aid=? and type=?", aid, 2).Count()

		if err != nil {
			log.Println("215")

			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// json.NewEncoder(w).Encode(pi_count)
		// json.NewEncoder(w).Encode(po_count)

		if pi_count <= po_count {

			// punchtable.Time = time.Now()
			punchtable.Time = time.Now().Add(time.Hour*5 + time.Minute*30)

			punchtable.Type = 1
			_, err := db.Model(&punchtable).Insert()

			if err != nil {
				log.Println("216")

				log.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

		} else {

			json.NewEncoder(w).Encode("You have already punch in")
			return

		}

	}

	json.NewEncoder(w).Encode("punch in successful")
}

func TeacherEntryPunchout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	teacherattendance := bean.TeacherAttendance{}

	_ = json.NewDecoder(r.Body).Decode(&teacherattendance)

	db := dataBase.Connect()

	defer db.Close()

	err := db.Model(&teacherattendance).Where("tid=? and date=? and month=? and year=? ", teacherattendance.Tid, time.Now().Day(), int(time.Now().Month()), time.Now().Year()).Select() // add date in where claise
	if err == pg.ErrNoRows {
		json.NewEncoder(w).Encode(" no data found  so go for punch in first")
		return

	} else if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return

	} else {

		aid := teacherattendance.Aid

		punchtable := bean.TeacherLogPunchs{Aid: aid}

		pi_count, err := db.Model(&punchtable).Where("aid=? and type=?", aid, 1).Count()

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		po_count, err := db.Model(&punchtable).Where("aid=? and type=?", aid, 2).Count()

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if pi_count > po_count {

			// punchtable.Time = time.Now()
			punchtable.Time = time.Now().Add(time.Hour*5 + time.Minute*30)

			punchtable.Type = 2
			_, err := db.Model(&punchtable).Insert()

			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

		} else {

			json.NewEncoder(w).Encode("You have already punch out")
			return

		}

	}

	json.NewEncoder(w).Encode("punch out successful")
}

func GetTeacherattendance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := dataBase.Connect()
	defer db.Close()
	var teacherattendance bean.TeacherAttendance
	json.NewDecoder(r.Body).Decode(&teacherattendance)
	var teacherattendancedetail []bean.TeacherAttendancetemp
	err := db.Model(&teacherattendancedetail).
		ColumnExpr(" DISTINCT teacher_attendances.date").
		Column("teacher_attendances.month").
		Column("teacher_attendances.year").
		Column("teacher_log_punchs.time").
		Column("teacher_log_punchs.type").
		Join("inner join teacher_attendances on teacher_attendances.aid=teacher_log_punchs.aid").
		Table("teacher_log_punchs").
		Where("teacher_attendances.tid=? AND teacher_attendances.month=? AND teacher_attendances.year=?", teacherattendance.Tid, teacherattendance.Month, teacherattendance.Year).
		Select()

	if err == pg.ErrNoRows {
		json.NewEncoder(w).Encode("no data found with this details ")
		return

	} else if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return

	} else {
		// fmt.Println("Third case")
		if teacherattendancedetail == nil {
			json.NewEncoder(w).Encode("teacher with this details doesn't exist")

		}
		json.NewEncoder(w).Encode(teacherattendancedetail)
		return
	}

}

func GetClassattendance(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get attendance request received")
	w.Header().Set("Content-Type", "application/json")
	db := dataBase.Connect()
	defer db.Close()
	// var student bean.Student

	var classtemp = &bean.Classtemp{}

	var classdata []bean.ClasstempRes
	json.NewDecoder(r.Body).Decode(&classtemp)

	// err := db.Model(&student).Where("class =?", classtemp.Class).Select()
	err := db.Model(&classdata).
		ColumnExpr(" DISTINCT students.sid").
		Column("students.name", "students.class", "student_attendances.date", "student_attendances.month", "student_attendances.year").
		Join("INNER JOIN student_attendances on student_attendances.sid=students.sid").
		Table("students").
		Where("student_attendances.date=? AND student_attendances.month=? AND student_attendances.year=?", classtemp.Date, classtemp.Month, classtemp.Year).
		Where("students.class =?", classtemp.Class).
		Select()
	// query := `select students.sid,students.name,students.class,student_attendances.date,student_attendances.month,student_attendances.year
	// 		from students inner join student_attendances on students.sid=student_attendances.sid where student_attendances.date=? and
	// 		student_attendances.month=? and student_attendances.year=? and students.class =?;`
	// _, err := db.Query(&classdata, query, classtemp.Date, classtemp.Month, classtemp.Year, classtemp.Class)
	if err == pg.ErrNoRows {
		json.NewEncoder(w).Encode("no students are in this class ")
		return

	} else if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return

	} else {

		json.NewEncoder(w).Encode(classdata)
		return
	}

}
